package main

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
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

func maxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n > 0 {
			max = findmax(n, max*n)
			min = findmin(n, min*n)
		} else {
			max, min = findmax(n, min*n), findmin(n, max*n)
		}
		if res < max {
			res = max
		}
	}

	return res
}

// @lc code=end
