package main

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}

	backtrack(n, 0, &res, []byte{})

	return res
}

func backtrack(n int, open int, res *[]string, tmp []byte) {
	if n == 0 && open == 0 {
		*res = append(*res, string(tmp))
	} else {
		if n > 0 {
			tmp = append(tmp, '(')
			backtrack(n-1, open+1, res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
		if open > 0 {
			tmp = append(tmp, ')')
			backtrack(n, open-1, res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
