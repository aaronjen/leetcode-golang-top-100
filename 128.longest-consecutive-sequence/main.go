package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	table := map[int]bool{}

	for _, n := range nums {
		table[n] = true
	}

	res := 0

	for key, value := range table {
		if value {
			tmp := extend(table, key)
			if tmp > res {
				res = tmp
			}
		}
	}

	return res
}

func extend(table map[int]bool, n int) int {
	left := n - 1
	for table[left] {
		table[left] = false
		left--
	}
	right := n + 1
	for table[right] {
		table[right] = false
		right++
	}

	return right - left - 1
}

// @lc code=end
