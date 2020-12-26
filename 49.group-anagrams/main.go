package main

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	table := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		keyArr := []byte(s)
		sort.Slice(keyArr, func(i, j int) bool {
			return keyArr[i] < keyArr[j]
		})
		key := string(keyArr)

		if _, ok := table[key]; ok {
			table[key] = append(table[key], s)
		} else {
			table[key] = []string{s}
		}
	}

	res := [][]string{}
	for _, val := range table {
		res = append(res, val)
	}

	return res
}

// @lc code=end
