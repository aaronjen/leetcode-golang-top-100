package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	backtrack(nums, &res, []int{}, used)

	return res
}

func backtrack(nums []int, res *[][]int, tmp []int, used []bool) {
	if len(tmp) == len(nums) {
		tt := make([]int, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	} else {
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			} else {
				tmp = append(tmp, nums[i])
				used[i] = true
				backtrack(nums, res, tmp, used)
				used[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
}

// @lc code=end
