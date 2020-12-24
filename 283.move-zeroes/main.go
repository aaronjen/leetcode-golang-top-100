package main

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			if cnt != i {
				nums[i] = 0
			}
			cnt++
		}
	}
}

// @lc code=end
