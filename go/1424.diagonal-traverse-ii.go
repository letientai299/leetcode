package main

// Diagonal Traverse II
//
// Medium
//
// https://leetcode.com/problems/diagonal-traverse-ii/
//
// Given a list of lists of integers, `nums`, return all elements of `nums` in
// diagonal order as shown in the below images.
//
// **Example 1:**
//
// **![](https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png)**
//
// ```
// Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,4,2,7,5,3,8,6,9]
//
// ```
//
// **Example 2:**
//
// **![](https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png)**
//
// ```
// Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
// Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
// Output: [1,4,2,5,3,8,6,9,7,10,11]
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [[1,2,3,4,5,6]]
// Output: [1,2,3,4,5,6]
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 10^5`
// - `1 <= nums[i].length <= 10^5`
// - `1 <= nums[i][j] <= 10^9`
// - There at most `10^5` elements in `nums`.
func findDiagonalOrder(nums [][]int) []int {
	m := len(nums)
	n := 0
	for _, row := range nums {
		if n < len(row) {
			n = len(row)
		}
	}

	// append
	lines := make([][]int, m+n-1)
	for y, row := range nums {
		for x, v := range row {
			lines[y+x] = append(lines[y+x], v)
		}
	}

	// flatten
	var res []int
	for _, line := range lines {
		for i := len(line) - 1; i >= 0; i-- {
			res = append(res, line[i])
		}
	}

	return res
}
