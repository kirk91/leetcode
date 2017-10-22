package main

import (
	"bytes"
	"math"
	"strconv"

	"github.com/kirk91/leetcode/assert"
)

func main() {
	assert.MustEqual("0.(4)", fractionToDecimal(4, 9))
}

func fractionToDecimal(n int, d int) string {
	if n == 0 {
		return "0"
	}

	var buffer bytes.Buffer
	if (n < 0 && d > 0) || (n > 0 && d < 0) {
		buffer.WriteString("-")
	}

	n = int(math.Abs(float64(n)))
	d = int(math.Abs(float64(d)))
	buffer.WriteString(strconv.Itoa(n / d))
	if n%d == 0 {
		return buffer.String()
	}

	buffer.WriteString(".")
	m := make(map[int]int)
	for r := n % d; r > 0; r %= d {
		if i, ok := m[r]; ok {
			buffer.WriteString(" ")
			b := buffer.Bytes()
			copy(b[i+1:], b[i:])
			b[i] = '('
			buffer.WriteString(")")
			break
		}

		m[r] = buffer.Len()
		r *= 10
		buffer.WriteString(strconv.Itoa(r / d))
	}
	return buffer.String()
}
