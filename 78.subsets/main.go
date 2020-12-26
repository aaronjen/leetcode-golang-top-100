package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}

	backtrack(nums, &res, []int{}, 0)

	return res
}

func backtrack(nums []int, res *[][]int, tmp []int, ind int) {
	tt := make([]int, len(tmp))
	copy(tt, tmp)
	*res = append(*res, tt)

	for i := ind; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backtrack(nums, res, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

// @lc code=end
