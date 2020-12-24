package main

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	res := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count++
		} else if res == nums[i] {
			count++
		} else {
			count--
		}
	}

	return res
}

// @lc code=end
