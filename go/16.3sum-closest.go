package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.71%)
 * Total Accepted:    431.4K
 * Total Submissions: 943.7K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */
func threeSumClosest(nums []int, t int) int {
	distance := func(a, b int) int {
		if a >= b {
			return a - b
		}

		return b - a
	}

	minDiff := math.MaxInt32
	var bestSum int
	var hasBest bool
	var sum int

	sort.Ints(nums)
	for i, x := range nums {
		for j, k := 0, len(nums)-1; j < k; {
			if i == j {
				j++
				continue
			}

			if k == i {
				k--
				continue
			}

			var diff int
			sum = x + nums[j] + nums[k]
			if !hasBest {
				bestSum = sum
				hasBest = true
				minDiff = distance(sum, t)
			} else {
				diff = distance(sum, t)
				if diff < minDiff {
					bestSum = sum
					minDiff = diff
				}
			}

			if bestSum == t {
				return bestSum
			}

			if sum < t {
				j++
			} else {
				k--
			}
		}
	}

	return bestSum
}
