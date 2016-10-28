// source: https://leetcode.com/problems/palindrome-number/

/*
Determine whether an integer is a palindrome. Do this without extra space.
*/
package main

import "log"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x > 0 {
		if x/div != x%10 {
			return false
		}
		x = x % div / 10
		div /= 100
	}
	return true
}

func main() {
	assert(isPalindrome(121), true)
	assert(isPalindrome(-121), false)
	assert(isPalindrome(1), true)
	assert(isPalindrome(12), false)
}

func assert(got, expected bool) {
	if got != expected {
		log.Fatalf("expected %v, but got %v", expected, got)
	}
}
