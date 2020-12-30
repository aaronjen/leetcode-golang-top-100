package main

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func pathSum(root *TreeNode, sum int) int {
	table := map[int]int{}

	res := 0

	helper(root, 0, sum, table, &res)

	return res
}

func helper(node *TreeNode, sum int, target int, table map[int]int, res *int) {
	if node == nil {
		return
	}

	sum += node.Val
	if sum == target {
		*res++
	}

	*res += table[sum-target]

	table[sum]++

	helper(node.Left, sum, target, table, res)
	helper(node.Right, sum, target, table, res)

	table[sum]--
}

// @lc code=end
