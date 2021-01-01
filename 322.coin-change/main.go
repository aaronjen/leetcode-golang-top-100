package main

import "math"

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32

		for _, c := range coins {
			if i >= c && dp[i-c] != math.MaxInt32 {
				dp[i] = findmin(dp[i-c]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// @lc code=end
