package main

import "math"

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	t1 := 0
	t2 := math.MinInt32
	t3 := 0
	for i := 0; i < len(prices); i++ {
		tmp := t3
		t3 = findmax(t3, t2+prices[i])
		t2 = findmax(t1-prices[i], t2)
		t1 = tmp
	}

	return findmax(t1, t3)
}

// @lc code=end
