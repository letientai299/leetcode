package main

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 *
 * https://leetcode.com/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Easy (30.23%)
 * Total Accepted:    70.7K
 * Total Submissions: 233.9K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * Given an integer array, you need to find one continuous subarray that if you
 * only sort this subarray in ascending order, then the whole array will be
 * sorted in ascending order, too.
 *
 * You need to find the shortest such subarray and output its length.
 *
 * Example 1:
 *
 * Input: [2, 6, 4, 8, 10, 9, 15]
 * Output: 5
 * Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make
 * the whole array sorted in ascending order.
 *
 *
 *
 * Note:
 *
 * Then length of the input array is in range [1, 10,000].
 * The input array may contain duplicates, so ascending order here means .
 *
 *
 */
func findUnsortedSubarray(nums []int) int {
	makeStack := func(start, end int, cmp func(a, b int) bool) int {
		maxCap := end - start + 1
		by := 1
		f := func(i int) bool {
			return i <= end
		}
		if end < start {
			by = -1
			maxCap = start - end + 1
			f = func(i int) bool {
				return i >= end
			}
		}

		stack := make([]int, 0, maxCap)
		n := len(stack)
		appendMode := true

		for i := start; f(i); i += by {
			x := nums[i]
			if appendMode {
				if n == 0 || cmp(x, stack[n-1]) {
					stack = append(stack, x)
					n++
				} else {
					appendMode = false
					for n > 0 && !cmp(x, stack[n-1]) {
						stack = stack[:n-1]
						n--
					}
					if n == 0 {
						break
					}
				}
			} else {
				for n > 0 && !cmp(x, stack[n-1]) {
					stack = stack[:n-1]
					n--
				}
				if n == 0 {
					break
				}
			}
		}
		return n
	}

	left := makeStack(0, len(nums)-1, func(a, b int) bool {
		return a >= b
	})

	if left >= len(nums) {
		return 0
	}

	right := makeStack(len(nums)-1, left, func(a, b int) bool {
		return a <= b
	})
	return len(nums) - left - right
}
