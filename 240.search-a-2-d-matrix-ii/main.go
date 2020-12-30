package main

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i := 0
	j := n - 1
	for i < m && j >= 0 {
		n := matrix[i][j]
		if n == target {
			return true
		}
		if n > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// @lc code=end
