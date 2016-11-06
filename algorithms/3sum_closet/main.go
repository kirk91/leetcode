// source: https://leetcode.com/problems/3sum-closest/

/*
Given an array S of n integers, find three integers in S such that the sum is closest to a given number,target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.
*/

package main

import (
	"sort"

	"github.com/kirk91/leetcode/assert"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)
	closetTarget := nums[0] + nums[1] + nums[n-1]
	var minDiff int
	if target > closetTarget {
		minDiff = target - closetTarget
	} else {
		minDiff = closetTarget - target
	}

	var left, right, diff, sum int
	for i := 0; i < n-1; i++ {
		left, right = i+1, n-1
		for left < right {
			sum = nums[left] + nums[right] + nums[i]
			if target > sum {
				diff = target - sum
				if diff < minDiff {
					minDiff = diff
					closetTarget = sum
				}
				left++
			} else {
				diff = sum - target
				if diff < minDiff {
					minDiff = diff
					closetTarget = sum
				}
				right--
			}
		}
	}

	return closetTarget
}

func main() {
	assert.MustEqual(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
