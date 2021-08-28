package main

// Trapping Rain Water
//
// Hard
//
// https://leetcode.com/problems/trapping-rain-water/
//
// Given `n` non-negative integers representing an elevation map where the width
// of each bar is `1`, compute how much water it can trap after raining.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)
//
// ```
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array
// [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
// are being trapped.
//
// ```
//
// **Example 2:**
//
// ```
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
// ```
//
// **Constraints:**
//
// - `n == height.length`
// - `1 <= n <= 2 * 104`
// - `0 <= height[i] <= 105`
func trap(height []int) int {
	n := len(height)
	l, r := 0, n-1

	all := 0
	pre := 0
	for l <= r {
		if height[l] == height[r] {
			all += (height[l] - pre) * (r - l + 1)
			pre = height[l]
			for h := height[l]; l <= r && height[l] <= h; l++ {
				all -= height[l]
			}
			for h := height[r]; l <= r && height[r] <= h; r-- {
				all -= height[r]
			}
			continue
		}

		if height[l] < height[r] {
			all += (height[l] - pre) * (r - l + 1)
			pre = height[l]
			for h := height[l]; l <= r && height[l] <= h; l++ {
				all -= height[l]
			}
			continue
		}

		all += (height[r] - pre) * (r - l + 1)
		pre = height[r]
		for h := height[r]; l <= r && height[r] <= h; r-- {
			all -= height[r]
		}
	}

	return all
}
