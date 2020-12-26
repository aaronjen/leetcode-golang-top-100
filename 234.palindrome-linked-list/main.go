package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	pad := &ListNode{}
	pad.Next = head
	fast := pad
	slow := pad
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast == slow {
		return true
	}

	var prev *ListNode
	mid := slow.Next
	slow.Next = nil
	for mid.Next != nil {
		tmp := mid.Next
		mid.Next = prev
		prev = mid
		mid = tmp
	}
	mid.Next = prev

	cur := head
	for cur != nil && mid != nil {
		if cur.Val != mid.Val {
			return false
		}
		cur = cur.Next
		mid = mid.Next
	}

	return true
}

// @lc code=end
