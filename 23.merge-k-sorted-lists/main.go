package main

import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := minHeap{}

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}

	res := &ListNode{}
	cur := res
	for len(h) != 0 {
		tmp := heap.Pop(&h).(*ListNode)
		cur.Next = tmp
		cur = cur.Next
		if tmp.Next != nil {
			heap.Push(&h, tmp.Next)
		}
	}

	return res.Next
}

// @lc code=end
