package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

var table []string = []string{
	"0",
	"1",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := []string{}

	backtrack(digits, 0, &res, []byte{})

	return res
}

func backtrack(digits string, ind int, res *[]string, tmp []byte) {
	if ind == len(digits) {
		*res = append(*res, string(tmp))
	} else {
		btn := int(digits[ind] - '0')
		s := table[btn]
		for i := 0; i < len(s); i++ {
			tmp = append(tmp, s[i])
			backtrack(digits, ind+1, res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
