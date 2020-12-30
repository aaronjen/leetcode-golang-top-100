package main

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	degrees := make([]int, numCourses)
	next := map[int][]int{}

	for i := 0; i < len(prerequisites); i++ {
		course, pre := prerequisites[i][0], prerequisites[i][1]

		degrees[course]++
		if _, ok := next[pre]; ok {
			next[pre] = append(next[pre], course)
		} else {
			next[pre] = []int{course}
		}
	}

	canTake := []int{}
	for i := 0; i < numCourses; i++ {
		if degrees[i] == 0 {
			canTake = append(canTake, i)
		}
	}

	numCourses -= len(canTake)

	for len(canTake) > 0 {
		tmp := []int{}
		for _, course := range canTake {
			nextCourses := next[course]
			for _, nc := range nextCourses {
				degrees[nc]--
				if degrees[nc] == 0 {
					tmp = append(tmp, nc)
					numCourses--
				}
			}
		}
		canTake = tmp
	}

	return numCourses == 0
}

// @lc code=end
