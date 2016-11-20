//source: https://leetcode.com/problems/valid-parentheses/

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

package main

import (
	"github.com/kirk91/leetcode/assert"
)

func isValid(s string) bool {
    var stack []byte
    for i, n := 0, len(s); i < n; i++{
        if isLeft(s[i]){
            stack = append(stack, s[i])
            continue
        }
        if len(stack) == 0 || !isClose(stack[len(stack)-1], s[i]){
            return false
        }
        stack = stack[:len(stack)-1]
    }
    return len(stack) == 0
}

func isLeft(b byte) bool{
    return b == '(' || b== '{' || b == '['
}

func isClose(l, r byte) bool {
    if l == '(' {
        return r == ')'
    }else if l == '{' {
        return r == '}'
    }else if l == '['{
        return r == ']'
    }
    return false
}

func main() {
	assert.MustEqual(true, isValid("{}()[]"))
	assert.MustEqual(true, isValid("{()}[]"))
	assert.MustEqual(false, isValid("{}()[]{"))
	assert.MustEqual(false, isValid("{}(){[]"))
	assert.MustEqual(false, isValid("{}()1{[]"))
	assert.MustEqual(false, isValid("{}(1{[]"))
}
