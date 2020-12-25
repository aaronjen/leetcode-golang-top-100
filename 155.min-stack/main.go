package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start

// MinStack type
type MinStack struct {
	allStack []int
	minStack []int
}

// Constructor func
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		allStack: []int{},
		minStack: []int{},
	}
}

// Push func
func (ms *MinStack) Push(x int) {
	ms.allStack = append(ms.allStack, x)
	if len(ms.minStack) == 0 {
		ms.minStack = append(ms.minStack, x)
	} else {
		min := ms.minStack[len(ms.minStack)-1]
		if x <= min {
			ms.minStack = append(ms.minStack, x)
		}
	}
}

// Pop func
func (ms *MinStack) Pop() {
	a := ms.allStack
	last := a[len(a)-1]
	ms.allStack = a[:len(a)-1]
	m := ms.minStack

	if last == m[len(m)-1] {
		ms.minStack = m[:len(m)-1]
	}
}

// Top func
func (ms *MinStack) Top() int {
	return ms.allStack[len(ms.allStack)-1]
}

// GetMin func
func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
