package main

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	var pre *int
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && cur.Val <= *pre {
			return false
		}
		if pre == nil {
			tmp := 0
			pre = &tmp
		}
		*pre = cur.Val
		if cur.Right != nil {
			root = cur.Right
		}
	}
	return true
}

// @lc code=end
