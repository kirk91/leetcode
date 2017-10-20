package main

import "github.com/kirk91/leetcode/assert"

func main() {
	assert.MustEqual(27, titleToNumber("AA"))
	assert.MustEqual(677, titleToNumber("ZA"))
}

func titleToNumber(s string) int {
	var num int
	for _, c := range s {
		num = num*26 + (c - 'A' + 1)
	}
	return num
}
