package main

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	org := head
	for org != nil {
		tmp := org.Next
		org.Next = &Node{
			Val:  org.Val,
			Next: tmp,
		}
		org = tmp
	}

	org = head
	for org != nil {
		if org.Random != nil {
			dup := org.Next
			dup.Random = org.Random.Next
		}
		org = org.Next.Next
	}

	org = head
	res := head.Next
	dup := head.Next

	for org != nil {
		org.Next, org = org.Next.Next, org.Next.Next
		if org != nil {
			dup.Next, dup = dup.Next.Next, dup.Next.Next
		}
	}

	return res
}

// @lc code=end
