// source: https://leetcode.com/problems/integer-to-roman/

/*
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

import (
	"log"
	"strings"
)

func intToRoman(num int) string {
	if num > 3999 {
		return ""
	}

	romans := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	}

	result := make([]string, 4)
	division := 1000
	for i := 0; i < 4; i++ {
		result[i] = romans[i][num/division]
		num %= division
		division /= 10
	}
	return strings.Join(result, "")
}

func main() {
	assert(intToRoman(1), "I")
	assert(intToRoman(2016), "MMXVI")
}

func assert(got, expected string) {
	if got != expected {
		log.Fatalf("expected %s, but got %s", expected, got)
	}
}
