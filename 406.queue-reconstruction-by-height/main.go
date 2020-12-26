package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func insert(s *[][]int, e []int, ind int) {
	old := *s
	old = append(old, e)
	copy(old[ind+1:], old[ind:])
	old[ind] = e
	*s = old
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][1] != people[j][1] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	res := [][]int{}
	for i := 0; i < len(people); i++ {
		p := people[i]
		cnt := p[1]
		ind := 0
		for cnt != 0 {
			if res[ind][0] >= p[0] {
				cnt--
			}
			ind++
		}
		insert(&res, p, ind)
	}

	return res
}

// @lc code=end
