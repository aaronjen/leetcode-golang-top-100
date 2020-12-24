package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	payload := 0
	for l1 != nil || l2 != nil {
		sum := payload
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		payload = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
	}

	if payload != 0 {
		cur.Next = &ListNode{
			Val: payload,
		}
	}

	return res.Next
}

// @lc code=end
