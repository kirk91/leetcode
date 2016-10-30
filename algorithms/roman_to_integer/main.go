// source: https://leetcode.com/problems/roman-to-integer/

/*
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

import "log"

func romanToInt(s string) int {
	roman := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	curVal := 0
	lastVal := roman[s[0]]
	sum := lastVal
	for i := 1; i < len(s); i++ {
		curVal = roman[s[i]]
		if curVal > lastVal {
			sum = sum + curVal - 2*lastVal
		} else {
			sum += curVal
		}
		lastVal = curVal
	}
	return sum
}

func main() {
	assert(romanToInt("DCXXI"), 621)
	assert(romanToInt("CDXXI"), 421)
}

func assert(got, expected int) {
	if got != expected {
		log.Fatalf("expected %d, but got %d", expected, got)
	}
}
