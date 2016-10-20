// source: https://leetcode.com/problems/longest-palindromic-substring/

/*
Given a string S, find the longest palindromic substring in S.
You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
*/

package main

import (
	"log"
)

func longestPalindrome(s string) string {
	var (
		n          = len(s)
		start, end int
	)
	for i := 0; i < n; i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := 0
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}

		if len > (end - start + 1) {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, start, end int) int {
	n := len(s)
	l, r := start, end
	for l >= 0 && r < n && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func main() {
	assert(longestPalindrome("abc"), "a")
	assert(longestPalindrome("abba"), "abba")
	assert(longestPalindrome("adbcbdx"), "dbcbd")
}

func assert(got, expected string) {
	if got != expected {
		log.Fatalf("expected: %s, got: %s", expected, got)
	}
}
