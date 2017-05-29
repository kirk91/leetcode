// source: https://leetcode.com/submissions/detail/82227154/

/*
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.
*/

package main

import (
	"sort"

	"github.com/kirk91/leetcode/assert"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeValues := threeSum(nums[i+1:], target-nums[i])
		for _, v := range threeValues {
			result = append(result, append(v, nums[i]))
		}
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
	var (
		result      [][]int
		n           = len(nums)
		left, right int
		cur, sum    int
	)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right = i+1, n-1
		cur = target - nums[i]
		for left < right {
			sum = nums[left] + nums[right]
			if sum == cur {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if sum > cur {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func main() {
	assert.MustEqual([][]int{[]int{-1, 1, 2, -2}, []int{0, 0, 2, -2}, []int{0, 0, 1, -1}}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
