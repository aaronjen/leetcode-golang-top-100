package main

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	beg := -1
	end := -2
	max := nums[0]
	min := nums[n-1]
	for i := 0; i < n; i++ {
		max = findmax(max, nums[i])
		min = findmin(min, nums[n-1-i])

		if nums[i] < max {
			end = i
		}
		if nums[n-1-i] > min {
			beg = n - 1 - i
		}
	}

	return end - beg + 1
}

// @lc code=end
