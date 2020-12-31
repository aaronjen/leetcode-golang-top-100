package main

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	p1 := 0
	p2 := nums[0]

	for i := 1; i < len(nums); i++ {
		p1, p2 = p2, findmax(p2, p1+nums[i])
	}

	return p2
}

// @lc code=end
