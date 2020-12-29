package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1

	for i := 0; i <= right; i++ {
		for nums[i] == 2 && i < right {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
		for nums[i] == 0 && i > left {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}

// @lc code=end
