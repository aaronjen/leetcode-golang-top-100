package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []*TreeNode{}
		row := []int{}
		for i := 0; i < len(queue); i++ {
			nd := queue[i]
			row = append(row, nd.Val)
			if nd.Left != nil {
				tmp = append(tmp, nd.Left)
			}
			if nd.Right != nil {
				tmp = append(tmp, nd.Right)
			}

		}
		queue = tmp
		res = append(res, row)
	}
	return res
}

// @lc code=end
