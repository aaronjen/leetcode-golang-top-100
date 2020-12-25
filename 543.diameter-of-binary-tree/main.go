package main

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	max(root, &res)

	return res
}

func max(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}

	left := max(node.Left, res)
	right := max(node.Right, res)

	if left+right > *res {
		*res = left + right
	}

	return findmax(left, right) + 1
}

// @lc code=end
