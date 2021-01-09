package main

/*
 * @lc app=leetcode id=845 lang=golang
 *
 * [845] Longest Mountain in Array
 *
 * https://leetcode.com/problems/longest-mountain-in-array/description/
 *
 * algorithms
 * Medium (35.52%)
 * Total Accepted:    61K
 * Total Submissions: 158.7K
 * Testcase Example:  '[2,1,4,7,3,2,5]'
 *
 * You may recall that an array arr is a mountain array if and only if:
 *
 *
 * arr.length >= 3
 * There exists some index i (0-indexed) with 0 < i < arr.length - 1 such
 * that:
 *
 * arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
 * arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 * Given an integer array arr, return the length of the longest subarray, which
 * is a mountain. Return 0 if there is no mountain subarray.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [2,1,4,7,3,2,5]
 * Output: 5
 * Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [2,2,2]
 * Output: 0
 * Explanation: There is no mountain.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 0 <= arr[i] <= 10^4
 *
 *
 *
 * Follow up:
 *
 *
 * Can you solve it using only one pass?
 * Can you solve it in O(1) space?
 *
 *
 */

// TODO (tai): do it again, can be simpler and faster
func longestMountain(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	best := 0
	cur := 1
	up := arr[0] < arr[1]
	isMountain := false
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			if isMountain {
				best = max(best, cur)
			}
			cur = 1
			isMountain = false
			up = false
			continue
		}

		if arr[i] > arr[i-1] {
			if !up {
				up = true
				if isMountain {
					best = max(best, cur)
				}
				isMountain = false
				cur = 2
			} else {
				cur++
			}
			continue
		}

		// arr[i] < arr[i-1]
		if up {
			up = false
			isMountain = true
		}

		cur++
	}

	if isMountain {
		best = max(best, cur)
	}
	if best < 3 {
		return 0
	}
	return best
}
