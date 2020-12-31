package main

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head.Next

	for fast != slow {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			return nil
		}
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// @lc code=end
