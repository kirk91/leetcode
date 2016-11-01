// source: https://leetcode.com/problems/3sum/

/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.
*/

package main

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	return threeSum(nums, 0)
}

func threeSum(nums []int, target int) [][]int {
	var sums [][]int
	if len(nums) < 3 {
		return sums
	}

	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		curTarget := target - nums[i]
		left, right := i+1, n-1
		for left < right {
			curSum := nums[left] + nums[right]
			if curSum == curTarget {
				sums = append(sums, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if curSum < curTarget {
				left++
			} else {
				right--
			}
		}
	}
	return sums
}

func main() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}
