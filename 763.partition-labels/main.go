package main

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}

	table := make([]int, 26)
	for i := 0; i < len(S); i++ {
		table[int(S[i]-'a')] = i
	}

	start := 0
	last := 0
	res := []int{}
	for i := 0; i < len(S); i++ {
		last = findmax(last, table[int(S[i]-'a')])
		if last == i {
			res = append(res, last-start+1)
			start = last + 1
		}
	}
	return res
}

// @lc code=end
