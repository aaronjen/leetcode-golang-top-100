package main

import "strings"

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func decodeString(s string) string {
	numStack := []int{}
	stringStack := []string{}
	curString := ""
	curNum := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '[' {
			numStack = append(numStack, curNum)
			stringStack = append(stringStack, curString)
			curString = ""
			curNum = 0
		} else if b == ']' {
			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			prevNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			var b strings.Builder
			b.WriteString(prevString)
			for i := 0; i < prevNum; i++ {
				b.WriteString(curString)
			}
			curString = b.String()
		} else if isNum(b) {
			curNum = curNum*10 + int(b-'0')
		} else {
			curString += string(b)
		}
	}
	return curString
}

// @lc code=end
