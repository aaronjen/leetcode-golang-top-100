package main

import "container/heap"

/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MedianFinder struct
type MedianFinder struct {
	left  maxHeap
	right minHeap
	even  bool
}

// Constructor func
/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		left:  maxHeap{},
		right: minHeap{},
		even:  true,
	}
}

// AddNum func
func (finder *MedianFinder) AddNum(num int) {
	if finder.even {
		heap.Push(&finder.right, num)
		heap.Push(&finder.left, heap.Pop(&finder.right))
	} else {
		heap.Push(&finder.left, num)
		heap.Push(&finder.right, heap.Pop(&finder.left))
	}
	finder.even = !finder.even
}

// FindMedian func
func (finder *MedianFinder) FindMedian() float64 {
	if finder.even {
		return float64(finder.left[0]+finder.right[0]) / 2
	}
	return float64(finder.left[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
