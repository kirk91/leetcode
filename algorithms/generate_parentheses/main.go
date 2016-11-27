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
	var res []string
	gen(n, n, make([]byte, 0, 2*n), &res)
	return res
}

func gen(left, right int, chs []byte, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(chs))
		return
	}
	if left > 0 {
		gen(left-1, right, append(chs, '('), res)
	}
	if right > left {
		gen(left, right-1, append(chs, ')'), res)
	}
}
func main() {
	assert.MustEqual([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
}
