package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	prev := 1
	for i := 1; i < len(nums); i++ {
		ind := len(nums) - i - 1

		prev *= nums[ind+1]
		res[ind] *= prev
	}

	return res
}

// @lc code=end
