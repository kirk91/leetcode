// source: https://leetcode.com/problems/generate-parentheses/

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

import "github.com/kirk91/leetcode/assert"

func generateParenthesis(n int) []string {
	var result []string
	gen([]byte{}, n, []byte{}, &result)
	return result
}

func gen(stack []byte, n int, chs []byte, result *[]string) {
	if len(stack) > 0 || n > 0 {
		if n == 0 {
			for i := 0; i < len(stack); i++ {
				chs = append(chs, ')')
			}
			stack = stack[:0]
		} else {
			if len(stack) == 0 {
				gen(append(stack, '('), n-1, append(chs, '('), result)
			} else {
				gen(append(stack, '('), n-1, append(chs, '('), result)
				gen(stack[:len(stack)-1], n, append(chs, ')'), result)
			}
		}
	}

	if len(stack) == 0 && n == 0 {
		*result = append(*result, string(chs))
	}
}

func main() {
	assert.MustEqual([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
}
