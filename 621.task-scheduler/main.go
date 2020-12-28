package main

/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	table := make([]int, 26)
	max := 0
	maxCount := 0
	for i := 0; i < len(tasks); i++ {
		t := int(tasks[i] - 'A')
		table[t]++
		if table[t] == max {
			maxCount++
		} else if table[t] > max {
			max = table[t]
			maxCount = 1
		}
	}
	res := (max-1)*(n+1) + maxCount
	if res < len(tasks) {
		return len(tasks)
	}
	return res
}

// @lc code=end
