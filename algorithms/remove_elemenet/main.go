// source: https://leetcode.com/problems/remove-element/

/*
Given an array and a value, remove all instances of that value in place and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example:
Given input array nums = [3,2,2,3], val = 3

Your function should return length = 2, with the first two elements of nums being 2.
*/
package main

import "github.com/kirk91/leetcode/assert"

func removeElement(nums []int, val int) int {
	var (
		end = 0
		n   = len(nums)
	)
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[end] = nums[i]
			end++
		}
	}
	nums = nums[:end]
	return end
}

func main() {
	assert.MustEqual(2, removeElement([]int{3, 2, 2, 3}, 3))
}
