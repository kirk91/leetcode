package main

import "github.com/kirk91/leetcode/assert"

func main() {
	assert.MustEqual(
		permute([]int{1, 2, 3}),
		[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	)
}

func permute(nums []int) [][]int {
	var result [][]int
	permuteRecursion(nums, nil, &result)
	return result
}

func permuteRecursion(nums []int, tmp []int, result *[][]int) {
	if len(nums) == 1 {
		tmp = append(tmp, nums[0])
		*result = append(*result, tmp)
		return
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		tmp = append(tmp, nums[i])
		var subNums []int
		subNums = append(subNums, nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)
		permuteRecursion(subNums, append([]int{}, tmp...), result)
		tmp = tmp[:len(tmp)-1]
	}
}
