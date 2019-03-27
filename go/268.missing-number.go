package main

/*
* @lc app=leetcode id=268 lang=golang
*
* [268] Missing Number
*
* https://leetcode.com/problems/missing-number/description/
*
* algorithms
* Easy (47.78%)
* Total Accepted:    255.9K
* Total Submissions: 535.6K
* Testcase Example:  '[3,0,1]'
*
* Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
* find the one that is missing from the array.
*
* Example 1:
*
*
* Input: [3,0,1]
* Output: 2
*
*
* Example 2:
*
*
* Input: [9,6,4,2,3,5,7,0,1]
* Output: 8
*
*
* Note:
* Your algorithm should run in linear runtime complexity. Could you implement
* it using only constant extra space complexity?
 */
func missingNumber(nums []int) int {
	// n := len(nums)

	// bitset := make([]uint32, n/32+1)
	// for _, x := range nums {
	// // turn on bit at x
	// bucket := x / 32
	// indexInBucket := x % 32
	// bitset[bucket] |= 1 << uint32(indexInBucket)
	// }

	// for i, b := range bitset {
	// if b == math.MaxUint32 {
	// continue
	// }
	// bitIndex := uint32(0)
	// for (b >> bitIndex & 1) == 1 {
	// bitIndex++
	// }

	// return i*32 + int(bitIndex)
	// }

	s := len(nums)
	for i, x := range nums {
		s += i - x
	}

	return s
}
