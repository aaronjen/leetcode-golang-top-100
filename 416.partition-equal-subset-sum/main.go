package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	half := sum / 2

	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for j := half; j >= n; j-- {
			if dp[j-n] {
				dp[j] = true
			}
		}
	}

	return dp[half]
}

// @lc code=end
