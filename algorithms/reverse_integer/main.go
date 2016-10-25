// source: https://leetcode.com/problems/reverse-integer/

/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321
*/

package main

import "log"

func reverse(x int) int {
	r := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		r = r*10 + x%10
		x = x / 10
	}
	if r < 0 {
		// overflow
		return 0
	}
	return sign * r
}

func main() {
	assert(reverse(123), 321)
	assert(reverse(100), 1)
	assert(reverse(10), 1)
	// overflow
	assert(reverse(1234567891234567899), 0)
}

func assert(got, expected int) {
	if got != expected {
		log.Fatalf("expected %d, but got %d", expected, got)
	}
}
