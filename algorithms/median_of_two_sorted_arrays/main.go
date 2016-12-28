// source: https://leetcode.com/problems/median-of-two-sorted-arrays/

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

package main

import (
	"github.com/kirk91/leetcode/assert"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	for ; i < m; i++ {
		sorted = append(sorted, nums1[i])
	}
	for ; j < n; j++ {
		sorted = append(sorted, nums2[j])
	}

	if (m+n)%2 == 0 {
		idx := (m + n) / 2
		return float64((sorted[idx] + sorted[idx-1])) / 2
	}
	return float64(sorted[(m+n)/2])
}

func main() {
	assert.MustEqual(2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.MustEqual(2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
