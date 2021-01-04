package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	max := 0
	jumps := 0
	curEnd := 0
	for i := 0; i < len(nums)-1; i++ {
		max = findmax(max, i+nums[i])
		if curEnd == i {
			jumps++
			curEnd = max
		}
	}

	return jumps
}

// @lc code=end
