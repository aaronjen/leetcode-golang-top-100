package main

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	need := make([]int, 26)
	count := 0
	for i := 0; i < len(p); i++ {
		pos := int(p[i] - 'a')

		need[pos]++
		count++
	}

	res := []int{}
	for i := 0; i < len(p); i++ {
		pos := int(s[i] - 'a')
		need[pos]--
		if need[pos] < 0 {
			count++
		} else {
			count--
		}
	}
	if count == 0 {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		right := i
		rightPos := int(s[right] - 'a')

		need[rightPos]--
		if need[rightPos] < 0 {
			count++
		} else {
			count--
		}

		left := right - len(p)
		leftPos := int(s[left] - 'a')
		need[leftPos]++
		if need[leftPos] > 0 {
			count++
		} else {
			count--
		}

		if count == 0 {
			res = append(res, left+1)
		}
	}
	return res
}

// @lc code=end
