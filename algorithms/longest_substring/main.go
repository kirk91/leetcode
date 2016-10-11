// source: https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

package main

import (
	"log"
)

func lengthOfLongestSubstring(s string) int {
	length, maxSubLength := len(s), 0
	for i := 0; i < length; i++ {
		b := make(map[byte]int)
		b[s[i]] = i
		for j := i + 1; j < length; j++ {
			if _, ok := b[s[j]]; ok {
				break
			} else {
				b[s[j]] = j
			}
		}
		subLength := len(b)
		if subLength > maxSubLength {
			maxSubLength = subLength
		}
	}
	return maxSubLength
}

func lengthOfLongestSubstring1(s string) int {
	var ok bool
	n := len(s)
	i, j, max, d := 0, 0, 0, 0
	b := make(map[byte]int)
	for i < n && j < n {
		if _, ok = b[s[j]]; !ok {
			b[s[j]] = j
			j++
			if d = j - i; d > max {
				max = d
			}
		} else {
			delete(b, s[i])
			i++
		}
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	var (
		ok           bool
		i, j, max, d int
	)
	n := len(s)
	b := make(map[byte]int)
	for ; j < n; j++ {
		if _, ok = b[s[j]]; ok {
			if b[s[j]] > i {
				i = b[s[j]]
			}
		}
		if d = j - i + 1; d > max {
			max = d
		}
		b[s[j]] = j + 1
	}
	return max
}

func lengthOfLongestSubstring3(s string) int {
	// asume character of string is ascii(<128)
	var (
		index  [128]int
		max, d int
		n      = len(s)
	)
	for i, j := 0, 0; j < n; j++ {
		if index[s[j]] > i {
			i = index[s[j]]
		}
		if d = j - i + 1; d > max {
			max = d
		}
		index[s[j]] = j + 1
	}
	return max
}

func main() {
	testLengthOfLongestSubstring("abcabcbb", 3)
	testLengthOfLongestSubstring("bbbbb", 1)
	testLengthOfLongestSubstring("pwwkew", 3)
	testLengthOfLongestSubstring("c", 1)
	testLengthOfLongestSubstring("", 0)
	testLengthOfLongestSubstring("au", 2)

	testLengthOfLongestSubstring1("abcabcbb", 3)
	testLengthOfLongestSubstring1("bbbbb", 1)
	testLengthOfLongestSubstring1("pwwkew", 3)
	testLengthOfLongestSubstring1("c", 1)
	testLengthOfLongestSubstring1("", 0)
	testLengthOfLongestSubstring1("au", 2)

	testLengthOfLongestSubstring2("abcabcbb", 3)
	testLengthOfLongestSubstring2("bbbbb", 1)
	testLengthOfLongestSubstring2("pwwkew", 3)
	testLengthOfLongestSubstring2("c", 1)
	testLengthOfLongestSubstring2("", 0)
	testLengthOfLongestSubstring2("au", 2)
	testLengthOfLongestSubstring2("abba", 2)

	testLengthOfLongestSubstring3("abcabcbb", 3)
	testLengthOfLongestSubstring3("bbbbb", 1)
	testLengthOfLongestSubstring3("pwwkew", 3)
	testLengthOfLongestSubstring3("c", 1)
	testLengthOfLongestSubstring3("", 0)
	testLengthOfLongestSubstring3("au", 2)
	testLengthOfLongestSubstring3("abba", 2)
}

func testLengthOfLongestSubstring(s string, expect int) {
	if length := lengthOfLongestSubstring(s); length != expect {
		log.Fatalf("compute longest length substring of %s got %d, but expecte %d", s, length, expect)
	}
}

func testLengthOfLongestSubstring1(s string, expect int) {
	if length := lengthOfLongestSubstring1(s); length != expect {
		log.Fatalf("compute longest length substring1 of %s got %d, but expecte %d", s, length, expect)
	}
}

func testLengthOfLongestSubstring2(s string, expect int) {
	if length := lengthOfLongestSubstring2(s); length != expect {
		log.Fatalf("compute longest length substring2 of %s got %d, but expecte %d", s, length, expect)
	}
}

func testLengthOfLongestSubstring3(s string, expect int) {
	if length := lengthOfLongestSubstring3(s); length != expect {
		log.Fatalf("compute longest length substring3 of %s got %d, but expecte %d", s, length, expect)
	}
}
