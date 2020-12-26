package main

import "container/heap"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start

type element struct {
	key int
	val int
}

type maxHeap []element

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i].val > m[j].val }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(element))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*m = old
	return x
}

func topKFrequent(nums []int, k int) []int {
	table := map[int]int{}

	for _, n := range nums {
		table[n]++
	}

	m := maxHeap{}
	for key, value := range table {
		heap.Push(&m, element{
			key: key,
			val: value,
		})
	}

	res := []int{}
	for i := 0; i < k; i++ {
		ele := heap.Pop(&m).(element)
		res = append(res, ele.key)
	}
	return res
}

// @lc code=end
