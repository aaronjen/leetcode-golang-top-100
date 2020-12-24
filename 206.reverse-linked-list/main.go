package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := head
	var prev *ListNode

	for res.Next != nil {
		tmp := res.Next
		res.Next = prev
		prev = res
		res = tmp
	}
	res.Next = prev

	return res
}

// @lc code=end
