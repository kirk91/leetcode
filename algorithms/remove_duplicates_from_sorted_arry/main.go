// souce: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

/*
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

*/
package main

import "github.com/kirk91/leetcode/assert"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	n := len(nums)
	end := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[end] {
			end++
			nums[end] = nums[i]
		}
	}
	nums = nums[:end+1]
	return end + 1
}

func main() {
	assert.MustEqual(0, removeDuplicates([]int{}))
	assert.MustEqual(1, removeDuplicates([]int{1, 1}))
	assert.MustEqual(2, removeDuplicates([]int{1, 1, 2}))
	assert.MustEqual(3, removeDuplicates([]int{1, 2, 3, 3}))
}
