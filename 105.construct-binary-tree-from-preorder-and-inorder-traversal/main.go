package main

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func buildTree(preorder []int, inorder []int) *TreeNode {
	ind := 0
	return solve(preorder, inorder, &ind, 0, len(preorder))
}

func solve(preorder []int, inorder []int, ind *int, left int, right int) *TreeNode {
	if left == right {
		return nil
	}
	val := preorder[*ind]
	pivot := left
	for i := left; i < right; i++ {
		if inorder[i] == val {
			pivot = i
		}
	}
	*ind++

	return &TreeNode{
		Val:   val,
		Left:  solve(preorder, inorder, ind, left, pivot),
		Right: solve(preorder, inorder, ind, pivot+1, right),
	}
}

// @lc code=end
