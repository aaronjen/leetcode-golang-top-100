package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxLeft := 0
	maxRight := 0
	res := 0

	for left <= right {
		if maxLeft < maxRight {
			if maxLeft < height[left] {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if maxRight < height[right] {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}

// @lc code=end
