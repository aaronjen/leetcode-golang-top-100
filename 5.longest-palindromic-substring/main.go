package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		extend(s, i, i, &res)
		extend(s, i, i+1, &res)
	}

	return res
}

func extend(s string, left, right int, res *string) {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		left--
		right++
	}
	if right-left-1 > len(*res) {
		*res = s[left+1 : right]
	}
}

// @lc code=end
