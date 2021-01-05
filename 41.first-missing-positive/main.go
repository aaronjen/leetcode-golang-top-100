package main

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		num := nums[i]
		for num > 0 && num <= n && nums[i] != nums[num-1] {
			nums[i], nums[num-1] = nums[num-1], nums[i]
			num = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		num := nums[i]
		if num != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// @lc code=end
