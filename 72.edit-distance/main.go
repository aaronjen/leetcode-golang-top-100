package main

/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	cur := make([]int, n+1)

	for j := 0; j <= n; j++ {
		cur[j] = j
	}

	for i := 1; i <= m; i++ {
		pre := cur[0]
		cur[0] = i
		for j := 1; j <= n; j++ {
			tmp := cur[j]
			if word1[i-1] == word2[j-1] {
				cur[j] = pre
			} else {
				cur[j] = findmin(cur[j], findmin(cur[j-1], pre)) + 1
			}
			pre = tmp
		}
	}

	return cur[n]
}

// @lc code=end
