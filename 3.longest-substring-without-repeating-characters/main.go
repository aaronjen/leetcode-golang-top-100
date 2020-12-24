package main

import "strings"

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	start := 0
	res := 1
	for i := 1; i < len(s); i++ {
		tmp := s[start:i]
		ind := strings.IndexByte(tmp, s[i])
		if ind == -1 {
			l := i - start + 1
			if l > res {
				res = l
			}
		} else {
			start += ind + 1
		}
	}

	return res
}

// @lc code=end
