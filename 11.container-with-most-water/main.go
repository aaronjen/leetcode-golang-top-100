package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left < right {
		lh := height[left]
		rh := height[right]

		val := 0

		if lh > rh {
			val = rh * (right - left)
			right--
		} else {
			val = lh * (right - left)
			left++
		}

		if val > res {
			res = val
		}
	}

	return res
}

// @lc code=end
