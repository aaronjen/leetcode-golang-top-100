package main

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func bs(nums []int, target int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[left] < nums[right] {
			break
		}
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	pivot := left

	ind := bs(nums, target, 0, pivot-1)
	if ind != -1 {
		return ind
	}

	ind = bs(nums, target, pivot, len(nums)-1)

	return ind
}

// @lc code=end
