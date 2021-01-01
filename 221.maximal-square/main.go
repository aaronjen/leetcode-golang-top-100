package main

/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	pre := make([]int, n)
	cur := make([]int, n)

	res := 0

	for i := 0; i < n; i++ {
		pre[i] = int(matrix[0][i] - '0')
		if pre[i] > res {
			res = pre[i]
		}
	}

	for i := 1; i < m; i++ {
		cur[0] = int(matrix[i][0] - '0')
		if cur[0] > res {
			res = cur[0]
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				cur[j] = findmin(cur[j-1], findmin(pre[j-1], pre[j])) + 1
				if cur[j] > res {
					res = cur[j]
				}
			} else {
				cur[j] = 0
			}
		}
		copy(pre, cur)
	}

	return res * res
}

// @lc code=end
