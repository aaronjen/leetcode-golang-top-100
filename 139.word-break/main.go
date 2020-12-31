package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func contains(wordDict []string, s string) bool {
	for _, word := range wordDict {
		if word == s {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				tmp := s[j:i]
				if contains(wordDict, tmp) {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end
