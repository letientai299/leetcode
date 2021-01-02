package main

/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 *
 * https://leetcode.com/problems/single-element-in-a-sorted-array/description/
 *
 * algorithms
 * Medium (57.56%)
 * Total Accepted:    173.1K
 * Total Submissions: 298.9K
 * Testcase Example:  '[1,1,2,3,3,4,4,8,8]'
 *
 * You are given a sorted array consisting of only integers where every element
 * appears exactly twice, except for one element which appears exactly once.
 * Find this single element that appears only once.
 *
 * Follow up: Your solution should run in O(log n) time and O(1) space.
 *
 *
 * Example 1:
 * Input: nums = [1,1,2,3,3,4,4,8,8]
 * Output: 2
 * Example 2:
 * Input: nums = [3,3,7,7,10,11,11]
 * Output: 10
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^5
 *
 */
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n%2 != 1 {
		return -1
	}

	if n == 1 {
		return nums[0]
	}

	var find func(l, r int) (int, bool)
	find = func(l, r int) (int, bool) {
		if l == r-3 {
			if nums[l] != nums[l+1] && nums[l+1] != nums[l+2] {
				return nums[l+1], true
			}

			if nums[l] == nums[l+1] && nums[l+1] != nums[l+2] {
				return nums[l+2], true
			}

			if nums[l] != nums[l+1] && nums[l+1] == nums[l+2] {
				return nums[l], true
			}

			return -1, false
		}

		m := (l + r) / 2

		if nums[m] != nums[m-1] && nums[m] != nums[m+1] {
			return nums[m], true
		}

		ml := m + 1
		if (ml-l)%2 == 0 {
			ml++
		}

		vl, okl := find(l, ml)
		if okl && vl != nums[ml] {
			return vl, true
		}

		mr := m + 1
		if (r-mr)%2 == 0 {
			mr--
		}

		vr, okr := find(mr, r)
		if okr && vr != nums[mr-1] {
			return vr, true
		}

		return -1, false
	}

	v, _ := find(0, len(nums))
	return v
}
