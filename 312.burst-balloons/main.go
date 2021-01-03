package main

/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */

// @lc code=start
func maxCoins(nums []int) int {
	validNums := make([]int, len(nums)+2)
	validNums[0] = 1
	n := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			validNums[n] = nums[i]
			n++
		}
	}
	validNums[n] = 1
	n++

	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}

	return burst(memo, validNums, 0, n-1)
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func burst(memo [][]int, nums []int, left, right int) int {
	if left+1 == right {
		return 0
	}

	if memo[left][right] > 0 {
		return memo[left][right]
	}

	res := 0
	for i := left + 1; i < right; i++ {
		res = findmax(res, nums[i]*nums[left]*nums[right]+burst(memo, nums, left, i)+burst(memo, nums, i, right))
	}
	memo[left][right] = res

	return res
}

// @lc code=end
