package main

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/kirk91/leetcode/assert"
)

func main() {
	assert.MustEqual("9534330", largestNumber([]int{3, 30, 34, 5, 9}))
	assert.MustEqual("0", largestNumber([]int{0, 0, 0}))
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(stringSlice(strs))
	for strs[0] == "0" && len(strs) > 1 {
		strs = strs[1:len(strs)]
	}
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

type stringSlice []string

func (p stringSlice) Len() int {
	return len(p)
}

func (p stringSlice) Less(i, j int) bool {
	return (p[i]+p[j] > p[j]+p[i])
}

func (p stringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
