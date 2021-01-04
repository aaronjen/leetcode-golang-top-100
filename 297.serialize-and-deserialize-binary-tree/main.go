package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

// Codec type
type Codec struct {
}

// Constructor func
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (t *Codec) serialize(root *TreeNode) string {
	s := buildString(root)
	return s
}

func buildString(node *TreeNode) string {
	if node == nil {
		return "x"
	}

	return strconv.Itoa(node.Val) +
		"," +
		buildString(node.Left) +
		"," +
		buildString(node.Right)
}

// Deserializes your encoded data to tree.
func (t *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	ind := 0
	return buildTree(arr, &ind)
}

func buildTree(data []string, ind *int) *TreeNode {
	if *ind >= len(data) {
		return nil
	}
	s := data[*ind]
	*ind++
	if s == "x" {
		return nil
	}
	val, _ := strconv.Atoi(s)
	return &TreeNode{
		Val:   val,
		Left:  buildTree(data, ind),
		Right: buildTree(data, ind),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
