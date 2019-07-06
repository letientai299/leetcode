package main

import (
	"math"
)

/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 *
 * https://leetcode.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (46.15%)
 * Total Accepted:    73.6K
 * Total Submissions: 159.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array, find three numbers whose product is maximum and
 * output the maximum product.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: 6
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: 24
 *
 *
 *
 *
 * Note:
 *
 *
 * The length of the given array will be in range [3,104] and all elements are
 * in the range [-1000, 1000].
 * Multiplication of any three numbers in the input won't exceed the range of
 * 32-bit signed integer.
 *
 *
 *
 *
 */
func maximumProduct(nums []int) int {
	// the max product 3 numbers in array must be either:
	// - product of 3 largest numbers
	// - product of the largest positive numbers and 2 smallest negative numbers

	// set of used indexes for the top largest values
	topIndexes := make(map[int]bool)
	bottomIndexes := make(map[int]bool)

	// findMaxMin returns the max number of that smaller than upper and the min
	// number that larger than lower
	findMaxMin := func(upper int, lower int) (up int, upId int, down int, downId int) {
		up = math.MinInt32
		upId = -1
		downId = -1
		down = math.MaxInt32
		for i, v := range nums {
			if up <= v && v <= upper && !topIndexes[i] {
				up = v
				upId = i
			}
			if down >= v && v >= lower && !bottomIndexes[i] {
				down = v
				downId = i
			}
		}

		return
	}

	max1, maxId1, min1, minId1 := findMaxMin(math.MaxInt32, math.MinInt32)
	topIndexes[maxId1] = true
	bottomIndexes[minId1] = true

	max2, maxId2, min2, minId2 := findMaxMin(max1, min1)
	topIndexes[maxId2] = true
	bottomIndexes[minId2] = true

	max3, _, _, _ := findMaxMin(max2, min2)

	res1 := max1 * max2 * max3
	res2 := max1 * min1 * min2

	if res1 > res2 {
		return res1
	}
	return res2
}
