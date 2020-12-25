package main

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	cur := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > cur {
			res = findmax(res, prices[i]-cur)
		} else {
			cur = prices[i]
		}
	}

	return res
}

// @lc code=end
