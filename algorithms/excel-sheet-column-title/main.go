package main

import "github.com/kirk91/leetcode/assert"

func main() {
	assert.MustEqual("ZA", convertToTitle(677))
	assert.MustEqual("A", convertToTitle(1))
	assert.MustEqual("AZ", convertToTitle(52))
}

func convertToTitle(n int) string {
	var chars []byte
	for n > 26 {
		r := n % 26
		if r == 0 {
			chars = append(chars, 'Z')
			n = n - 26
		} else {
			chars = append(chars, byte(r-1+'A'))
		}
		n = n / 26
	}
	chars = append(chars, byte(n-1+'A'))
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
