package main

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		next := nums[i]
		for next != 0 {
			next, nums[next-1] = nums[next-1], 0
		}
	}

	res := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res = append(res, i+1)
		}
	}

	return res
}

// @lc code=end
