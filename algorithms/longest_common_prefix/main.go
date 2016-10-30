// source: https://leetcode.com/problems/longest-common-prefix/

/*
Write a function to find the longest common prefix string amongst an array of strings.
*/

package main

import "log"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var (
		i, j      int
		base      = strs[0]
		baseLen   = len(base)
		strsCount = len(strs)
	)
loop:
	for i = 0; i < baseLen; i++ {
		for j = 1; j < strsCount; j++ {
			if i >= len(strs[j]) || strs[j][i] != base[i] {
				break loop
			}
		}
	}
	return base[:i]
}

func main() {
	assert(longestCommonPrefix([]string{}), "")
	assert(longestCommonPrefix([]string{"a", "b"}), "")
	assert(longestCommonPrefix([]string{"a", "ab"}), "a")
}

func assert(got, expected string) {
	if got != expected {
		log.Fatalf("expected %s, but got %s", expected, got)
	}
}
