package main

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	res1, _ := solve(root)
	return res1
}

func solve(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	left1, left2 := solve(node.Left)
	right1, right2 := solve(node.Right)

	return findmax(left1+right1, left2+right2+node.Val), left1 + right1
}

// @lc code=end
