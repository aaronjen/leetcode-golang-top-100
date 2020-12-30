package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	if S > 1000 {
		return 0
	}
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, 2001)
	}

	dp[0][1000] = 1

	for i := 0; i < len(nums); i++ {
		n := nums[i]

		for j := 0; j < 2001; j++ {
			if j+n < 2001 {
				dp[i+1][j+n] += dp[i][j]
			}
			if j-n >= 0 {
				dp[i+1][j-n] += dp[i][j]
			}
		}
	}

	return dp[len(nums)][1000+S]
}

// @lc code=end
