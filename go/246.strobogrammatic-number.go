package main

/*
 * @lc app=leetcode id=246 lang=golang
 *
 * [246] Strobogrammatic Number
 *
 * https://leetcode.com/problems/strobogrammatic-number/description/
 *
 * algorithms
 * Easy (43.53%)
 * Total Accepted:    65.4K
 * Total Submissions: 149.5K
 * Testcase Example:  '"69"'
 *
 * A strobogrammatic number is a number that looks the same when rotated 180
 * degrees (looked at upside down).
 *
 * Write a function to determine if a number is strobogrammatic. The number is
 * represented as a string.
 *
 * Example 1:
 *
 *
 * Input:  "69"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:  "88"
 * Output: true
 *
 * Example 3:
 *
 *
 * Input:  "962"
 * Output: false
 *
 */
func isStrobogrammatic(num string) bool {
	upsidedowns := map[uint8]uint8{
		6: 9,
		9: 6,
		8: 8,
		0: 0,
		1: 1,
	}

	n := len(num)
	for i := 0; i < n/2+1; i++ {
		x := num[i] - '0'
		if y, ok := upsidedowns[x]; !ok || num[n-i-1]-'0' != y {
			return false
		}
	}

	return true
}
