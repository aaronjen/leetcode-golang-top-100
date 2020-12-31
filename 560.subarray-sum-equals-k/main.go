package main

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	table := map[int]int{}
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		sum += n

		if sum == k {
			res++
		}

		res += table[sum-k]

		table[sum]++
	}

	return res
}

// @lc code=end
