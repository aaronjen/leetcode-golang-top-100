package main

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
type monotonicQueue []int

func (q *monotonicQueue) push(n int) {
	pq := *q
	ed := len(pq) - 1
	for ed >= 0 && n > pq[ed] {
		ed--
	}
	if ed < 0 {
		*q = []int{}
	} else {
		*q = pq[:ed+1]
	}
	*q = append(*q, n)
}

func (q *monotonicQueue) pop() {
	pq := *q
	if len(pq) > 1 {
		*q = pq[1:]
	} else {
		*q = []int{}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	q := monotonicQueue{}

	ans := []int{}
	for i := 0; i < len(nums); i++ {
		q.push(nums[i])
		if i-k+1 >= 0 {
			ans = append(ans, q[0])
			if nums[i-k+1] == q[0] {
				q.pop()
			}
		}
	}
	return ans
}

// @lc code=end
