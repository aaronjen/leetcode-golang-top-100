package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	prevStart, prevEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]

		if prevEnd >= curStart {
			prevEnd = findmax(prevEnd, curEnd)
		} else {
			res = append(res, []int{prevStart, prevEnd})
			prevStart = curStart
			prevEnd = curEnd
		}
	}

	res = append(res, []int{prevStart, prevEnd})

	return res

}

// @lc code=end
