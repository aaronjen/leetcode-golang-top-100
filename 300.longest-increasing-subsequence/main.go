package main

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		n := nums[i]

		pass := true
		for i := 0; i < len(tails); i++ {
			if tails[i] >= n {
				tails[i] = n
				pass = false
				break
			}
		}
		if pass {
			tails = append(tails, n)
		}
	}

	return len(tails)
}

// @lc code=end
