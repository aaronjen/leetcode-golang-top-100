package main

import "math"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	l := 0
	r := n1
	k := (n1 + n2 + 1) / 2
	for l < r {
		m1 := (l + r) / 2
		m2 := k - m1
		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}

	m1 := l
	m2 := k - m1

	m1num1 := math.MinInt32
	if m1 > 0 {
		m1num1 = nums1[m1-1]
	}

	m2num1 := math.MinInt32
	if m2 > 0 {
		m2num1 = nums2[m2-1]
	}

	c1 := float64(findmax(m1num1, m2num1))
	if (n1+n2)%2 != 0 {
		return c1
	}

	m1num2 := math.MaxInt32
	if m1 < n1 {
		m1num2 = nums1[m1]
	}

	m2num2 := math.MaxInt32
	if m2 < n2 {
		m2num2 = nums2[m2]
	}
	c2 := float64(findmin(m1num2, m2num2))

	return (c1 + c2) / 2
}

// @lc code=end
