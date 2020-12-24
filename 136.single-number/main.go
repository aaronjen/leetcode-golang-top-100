package main

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}

	return nums[len(nums)-1]
}

// @lc code=end
