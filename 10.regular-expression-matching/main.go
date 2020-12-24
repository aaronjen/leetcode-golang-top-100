package main

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			charS := s[i-1]
			charP := p[j-1]

			if charP == '.' || charS == charP {
				dp[i][j] = dp[i-1][j-1]
			} else if charP == '*' {
				prevP := p[j-2]

				if prevP == '.' || prevP == charS {
					// count as empty
					// count as single
					// count as multiple
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
