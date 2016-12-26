// source: https://leetcode.com/problems/divide-two-integers/

/*
Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.
*/

package main

import (
	"math"

	"github.com/kirk91/leetcode/assert"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt64
	}
	if dividend == 0 {
		return 0
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	var result int
	for dividend >= divisor {
		var shift uint
		for dividend >= divisor<<shift {
			shift++
		}
		dividend -= divisor << (shift - 1)
		result += 1 << (shift - 1)
	}

	if sign == -1 {
		return -result
	}
	return result
}

func main() {
	assert.MustEqual(math.MaxInt64, divide(10, 0))
	assert.MustEqual(3, divide(9, 3))
	assert.MustEqual(3, divide(10, 3))
	assert.MustEqual(4, divide(12, 3))
	assert.MustEqual(1, divide(-1, -1))
	assert.MustEqual(1, divide(-1, -1))
	assert.MustEqual(-5, divide(10, -2))
}
