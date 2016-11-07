// source: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

/*
ven a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.
*/

package main

import "github.com/kirk91/leetcode/assert"

var dict = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "qprs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combs := make([]string, 0, len(digits))
	dfs(digits, &combs, 0, "")
	return combs
}

func dfs(digits string, combsPointer *[]string, index int, comb string) {
	if index == len(digits) {
		*combsPointer = append(*combsPointer, comb)
		return
	}
	lett := dict[digits[index]-'0']
	for i := 0; i < len(lett); i++ {
		dfs(digits, combsPointer, index+1, comb+string(lett[i]))
	}
}

func main() {
	assert.MustEqual([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
}
