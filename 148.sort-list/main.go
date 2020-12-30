package main

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head.Next
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(mid)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
	l1.Next = merge(l1.Next, l2)
	return l1
}

// @lc code=end
