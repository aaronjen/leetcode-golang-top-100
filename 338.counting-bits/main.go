package main

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = res[i>>1] + (i & 1)
	}

	return res
}

// @lc code=end
