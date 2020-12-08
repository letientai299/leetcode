package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (26.94%)
 * Total Accepted:    226.3K
 * Total Submissions: 749.4K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non-negative integers nums, arrange them such that they form
 * the largest number.
 *
 * Note: The result may be very large, so you need to return a string instead
 * of an integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,2]
 * Output: "210"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,30,34,5,9]
 * Output: "9534330"
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1]
 * Output: "1"
 *
 *
 * Example 4:
 *
 *
 * Input: nums = [10]
 * Output: "10"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 10^9
 *
 *
 */
func largestNumber(nums []int) string {
	ss := make([]string, 0, len(nums))
	for _, n := range nums {
		s := strconv.Itoa(n)
		ss = append(ss, s)
	}

	var comp func(a, b string) bool
	comp = func(a, b string) bool {
		if a == "" {
			return true
		}

		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] > b[i] {
				return true
			} else if a[i] < b[i] {
				return false
			}
		}

		if len(a) == len(b) {
			return true
		}

		if len(a) > len(b) {
			a = a[len(b):]
		} else {
			b = b[len(a):]
		}

		return comp(a, b)
	}

	sort.Slice(ss, func(i, j int) bool {
		return comp(ss[i], ss[j])
	})

	s := strings.Join(ss, "")
	i := 0
	for i < len(s)-1 && s[i] == '0' {
		i++
	}

	return s[i:]
}
