package main

/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	heights := make([]int, n)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		res = findmax(res, maxHistogram(heights))
	}
	return res
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxHistogram(heights []int) int {
	lessFromLeft := make([]int, len(heights))
	lessFromRight := make([]int, len(heights))

	lessFromLeft[0] = -1
	lessFromRight[len(heights)-1] = len(heights)

	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	res := 0
	for i := 0; i < len(heights); i++ {
		res = findmax(res, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}

	return res
}

// @lc code=end
