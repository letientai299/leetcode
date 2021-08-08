package main

// Tuple with Same Product
//
// Medium
//
// https://leetcode.com/problems/tuple-with-same-product/
//
// Given an array `nums` of **distinct** positive integers, return _the number
// of tuples_`(a, b, c, d)` _such that_ `a * b = c * d` _where_ `a` _,_ `b` _,_
// `c` _, and_ `d` _are elements of_ `nums` _, and_ `a != b != c != d` _._
//
// **Example 1:**
//
// ```
// Input: nums = [2,3,4,6]
// Output: 8
// Explanation: There are 8 valid tuples:
// (2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
// (3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,2,4,5,10]
// Output: 16
// Explanation: There are 16 valids tuples:
// (1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
// (2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
// (2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
// (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [2,3,4,6,8,12]
// Output: 40
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [2,3,5,7]
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 1000`
// - `1 <= nums[i] <= 104`
// - All elements in `nums` are **distinct**.
func tupleSameProduct(nums []int) int {
	m := make(map[int]int, len(nums))
	n := len(nums)
	r := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := nums[i] * nums[j]
			r += m[v]
			m[v]++
		}
	}

	// shuffle the tuple
	return 8 * r
}

// 2 6 3 4
// 4 12 6 8
// 2 12 4 6
// 2 12 3 8
// 4 6 3 8
