// source: https://leetcode.com/problems/zigzag-conversion/

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR"
*/

package main

import (
	"log"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var (
		i, j   int
		n      = len(s)
		size   = 2*numRows - 2
		result = make([]byte, 0, n)
	)
	for i = 0; i < numRows; i++ {
		for j = i; j < n; j += size {
			result = append(result, s[j])
			if i != 0 && i != numRows-1 && (j+size-2*i) < n {
				result = append(result, s[j+size-2*i])
			}
		}
	}
	return string(result)
}

func main() {
	assert("PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert("ABCED", convert("ABCDE", 4))
}

func assert(expect, got string) {
	if expect != got {
		log.Fatalf("expect %s, but got: %s", expect, got)
	}
}
