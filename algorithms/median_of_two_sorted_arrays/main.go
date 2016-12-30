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

import "github.com/kirk91/leetcode/assert"

// O(m+n)
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

// O(logm + logn)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		return float64(findKth(nums1, nums2, n/2)+findKth(nums1, nums2, n/2+1)) / 2
	}
	return float64(findKth(nums1, nums2, (n+1)/2))
}

func findKth(a []int, b []int, k int) int {
	m, n := len(a), len(b)
	if m > n {
		return findKth(b, a, k)
	}
	if k > (m + n) {
		panic("index out of range")
	}
	if m == 0 {
		return b[k-1]
	}
	if n == 0 {
		return a[k-1]
	}

	if k == 1 {
		if a[0] < b[0] {
			return a[0]
		} else {
			return b[0]
		}
	}

	var i int
	if k/2 > m {
		i = m
	} else {
		i = k / 2
	}
	j := k - i
	if a[i-1] >= b[j-1] {
		return findKth(a, b[j:], k-j)
	}
	return findKth(a[i:], b, k-i)
}

// O(log(min(m,n))
func findMedianSortedArrays3(a, b []int) float64 {
	m, n := len(a), len(b)
	if m > n {
		a, b, m, n = b, a, n, m
	}
	l, r := 0, m
	halfLen := (m + n + 1) / 2

	var i, j int
	var median float64
	// binary search
	for l <= r {
		i = (l + r) / 2
		j = halfLen - i

		if i < m && b[j-1] > a[i] {
			l++
		} else if i > 0 && a[i-1] > b[j] {
			r--
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = b[j-1]
			} else if j == 0 {
				maxLeft = a[i-1]
			} else {
				maxLeft = max(a[i-1], b[j-1])
			}

			if (m+n)%2 == 1 {
				median = float64(maxLeft)
				break
			}

			var minRight int
			if i == m {
				minRight = b[j]
			} else if j == n {
				minRight = a[i]
			} else {
				minRight = min(a[i], b[j])
			}

			median = float64(maxLeft+minRight) / 2
			break
		}
	}
	return median
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	assert.MustEqual(2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.MustEqual(2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.MustEqual(2.0, findMedianSortedArrays2([]int{1, 3}, []int{2}))
	assert.MustEqual(2.5, findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
	assert.MustEqual(2.0, findMedianSortedArrays3([]int{1, 3}, []int{2}))
	assert.MustEqual(2.5, findMedianSortedArrays3([]int{1, 2}, []int{3, 4}))
	assert.MustEqual(2.5, findMedianSortedArrays3([]int{3, 4}, []int{1, 2}))
}
