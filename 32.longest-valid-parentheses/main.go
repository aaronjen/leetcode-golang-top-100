package main

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	n := len(s)
	open := 0
	cur := 0
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			open++
			cur++
		} else {
			open--
			cur++
			if open == 0 {
				res = findmax(res, cur)
			}
			if open < 0 {
				open = 0
				cur = 0
			}
		}
	}

	open = 0
	cur = 0

	for i := 0; i < n; i++ {
		if s[n-i-1] == ')' {
			open++
			cur++
		} else {
			open--
			cur++
			if open == 0 {
				res = findmax(res, cur)
			}
			if open < 0 {
				open = 0
				cur = 0
			}
		}
	}

	return res
}

// @lc code=end
