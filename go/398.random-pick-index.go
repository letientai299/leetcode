package main

import "math/rand"

/*
 * @lc app=leetcode id=398 lang=golang
 *
 * [398] Random Pick Index
 *
 * https://leetcode.com/problems/random-pick-index/description/
 *
 * algorithms
 * Medium (52.11%)
 * Total Accepted:    108K
 * Total Submissions: 188.1K
 * Testcase Example:  '["Solution","pick"]\n[[[1,2,3,3,3]],[3]]'
 *
 * Given an array of integers with possible duplicates, randomly output the
 * index of a given target number. You can assume that the given target number
 * must exist in the array.
 *
 * Note:
 * The array size can be very large. Solution that uses too much extra space
 * will not pass the judge.
 *
 * Example:
 *
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) should return either index 2, 3, or 4 randomly. Each index should
 * have equal probability of returning.
 * solution.pick(3);
 *
 * // pick(1) should return 0. Since in the array only nums[0] is equal to 1.
 * solution.pick(1);
 *
 *
 */
type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	if len(nums) == 0 {
		return Solution{m: m}
	}

	pre := nums[0]
	m[pre] = append(m[pre], 0)
	for i := 1; i < len(nums); i++ {
		if pre == nums[i] {
			continue
		}

		m[pre] = append(m[pre], i-1)
		pre = nums[i]
		m[pre] = append(m[pre], i)
	}

	m[pre] = append(m[pre], len(nums)-1)

	return Solution{m: m}
}

func (s *Solution) Pick(target int) int {
	if a, ok := s.m[target]; ok {
		i := rand.Int() % (a[1] - a[0] + 1)
		return i + a[0]
	} else {
		return -1
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
