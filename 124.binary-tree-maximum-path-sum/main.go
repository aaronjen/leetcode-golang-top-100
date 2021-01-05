package main

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

func maxPathSum(root *TreeNode) int {
	res := root.Val
	helper(root, &res)
	return res
}

func helper(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	sum := node.Val
	left := helper(node.Left, res)
	right := helper(node.Right, res)
	if left > 0 {
		sum += left
	}
	if right > 0 {
		sum += right
	}

	if sum > *res {
		*res = sum
	}

	return node.Val + findmax(0, findmax(left, right))
}

// @lc code=end
