// source: https://leetcode.com/problems/container-with-most-water/

/*
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.
*/

package main

import "log"

func maxArea(height []int) int {
	area, maxArea := 0, 0
	l, r := 0, len(height)-1
	w, h := 0, 0
	for l < r {
		w = r - l
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		area = w * h
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	assert(maxArea([]int{1, 1}), 1)
	assert(maxArea([]int{1, 3, 5}), 3)
}

func assert(got, expected int) {
	if got != expected {
		log.Fatalf("expected %d, but got %d", expected, got)
	}
}
