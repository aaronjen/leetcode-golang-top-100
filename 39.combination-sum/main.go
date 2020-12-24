package main

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	backtrack(candidates, target, &res, []int{}, 0)

	return res
}

func backtrack(candidates []int, target int, res *[][]int, tmp []int, ind int) {
	if target == 0 {
		tt := make([]int, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	} else {
		for i := ind; i < len(candidates); i++ {
			n := candidates[i]
			if n <= target {
				tmp = append(tmp, n)
				backtrack(candidates, target-n, res, tmp, i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
}

// @lc code=end
