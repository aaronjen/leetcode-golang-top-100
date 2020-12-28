package main

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	prev := root
	stack := []*TreeNode{}
	for {
		if prev.Right != nil {
			stack = append(stack, prev.Right)
		}
		if prev.Left != nil {
			stack = append(stack, prev.Left)
		}
		if len(stack) == 0 {
			break
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev.Right = tmp
		prev.Left = nil
		prev = tmp
	}
}

// @lc code=end
