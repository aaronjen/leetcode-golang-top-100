package main

import (
	"math"
)

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([]int, n)
	for j := 1; j < n; j++ {
		dp[j] = math.MaxInt32
	}

	for i := 0; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = findmin(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[n-1]
}

// @lc code=end
