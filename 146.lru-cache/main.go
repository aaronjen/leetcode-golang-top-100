package main

import (
	"container/list"
)

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

type pair struct {
	key int
	val int
}

// LRUCache type
type LRUCache struct {
	cap       int
	valueList *list.List
	table     map[int]*list.Element
}

// Constructor func
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:       capacity,
		valueList: list.New(),
		table:     map[int]*list.Element{},
	}
}

// Get func
func (cc *LRUCache) Get(key int) int {
	if ele, ok := cc.table[key]; ok {
		cc.valueList.MoveToFront(ele)
		p := ele.Value.(*pair)
		return p.val
	}
	return -1
}

// Put func
func (cc *LRUCache) Put(key int, value int) {
	l := cc.valueList

	if ele, ok := cc.table[key]; ok {
		p := ele.Value.(*pair)
		p.val = value
		l.MoveToFront(ele)
	} else {
		front := l.PushFront(&pair{
			key: key,
			val: value,
		})
		cc.table[key] = front
	}

	number := len(cc.table)
	if cc.cap < number {
		b := l.Back()
		p := b.Value.(*pair)
		delete(cc.table, p.key)
		l.Remove(b)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
