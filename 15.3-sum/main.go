package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		prevLeft := math.MaxInt32
		prevRight := math.MaxInt32
		for left < right {
			leftNum := nums[left]
			rightNum := nums[right]
			if prevLeft == leftNum {
				left++
				continue
			}
			if prevRight == rightNum {
				right--
				continue
			}

			sum := leftNum + rightNum + nums[i]

			if sum == 0 {
				res = append(res, []int{nums[i], leftNum, rightNum})
				prevLeft = leftNum
				prevRight = rightNum
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

// @lc code=end
