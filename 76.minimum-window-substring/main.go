package main

import (
	"math"
)

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	need := map[byte]int{}
	cnt := 0
	for _, r := range t {
		b := byte(r)
		if need[b] == 0 {
			cnt++
		}
		need[b]++
	}

	resLen := math.MaxInt32
	res := ""
	start := 0
	for i := 0; i < len(s); i++ {
		rb := s[i]
		need[rb]--
		if need[rb] == 0 {
			cnt--
		}
		for cnt == 0 {
			tmp := s[start : i+1]
			if len(tmp) < resLen {
				resLen = len(tmp)
				res = tmp
			}

			lb := s[start]
			need[lb]++
			if need[lb] == 1 {
				cnt++
			}
			start++
		}
	}

	return res
}

// @lc code=end
