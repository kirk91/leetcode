package main

import "github.com/kirk91/leetcode/assert"

func findKth(a []int, b []int, k int) int {
	m, n := len(a), len(b)
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

func main() {
	assert.MustEqual(0, findKth([]int{1, 3, 4}, []int{0, 2, 8, 9, 10}, 1))
	assert.MustEqual(2, findKth([]int{1, 3, 4}, []int{0, 2, 8, 9, 10}, 3))
	assert.MustEqual(4, findKth([]int{1, 3, 4}, []int{0, 2, 8, 9, 10}, 5))
	assert.MustEqual(10, findKth([]int{1, 3, 4}, []int{0, 2, 8, 9, 10}, 8))
	assert.MustEqual(11, findKth([]int{1, 3, 4, 11}, []int{0, 2, 8, 9, 10}, 9))
}
