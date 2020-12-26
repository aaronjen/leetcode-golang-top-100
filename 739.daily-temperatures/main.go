package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	bigger := make([]int, len(T))

	for i := 1; i < len(T); i++ {
		ind := len(T) - i - 1
		if T[ind+1] > T[ind] {
			res[ind] = 1
			bigger[ind] = ind + 1
			continue
		}
		next := bigger[ind+1]
		for next != 0 {
			if T[ind] < T[next] {
				res[ind] = next - ind
				bigger[ind] = next
				break
			}
			next = bigger[next]
		}
	}

	return res
}

// @lc code=end
