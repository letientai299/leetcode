package main

/*
 * @lc app=leetcode id=1073 lang=golang
 *
 * [1073] Adding Two Negabinary Numbers
 *
 * https://leetcode.com/problems/adding-two-negabinary-numbers/description/
 *
 * algorithms
 * Medium (32.58%)
 * Total Accepted:    8K
 * Total Submissions: 23.1K
 * Testcase Example:  '[1,1,1,1,1]\n[1,0,1]'
 *
 * Given two numbers arr1 and arr2 in base -2, return the result of adding them
 * together.
 *
 * Each number is given in array format:  as an array of 0s and 1s, from most
 * significant bit to least significant bit.  For example, arr = [1,1,0,1]
 * represents the number (-2)^3 + (-2)^2 + (-2)^0 = -3.  A number arr in array,
 * format is also guaranteed to have no leading zeros: either arr == [0] or
 * arr[0] == 1.
 *
 * Return the result of adding arr1 and arr2 in the same format: as an array of
 * 0s and 1s with no leading zeros.
 *
 *
 * Example 1:
 *
 *
 * Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
 * Output: [1,0,0,0,0]
 * Explanation: arr1 represents 11, arr2 represents 5, the output represents
 * 16.
 *
 *
 * Example 2:
 *
 *
 * Input: arr1 = [0], arr2 = [0]
 * Output: [0]
 *
 *
 * Example 3:
 *
 *
 * Input: arr1 = [0], arr2 = [1]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr1.length, arr2.length <= 1000
 * arr1[i] and arr2[i] are 0 or 1
 * arr1 and arr2 have no leading zeros
 *
 *
 */

func addNegabinary(a []int, b []int) []int {
	if len(a) < len(b) {
		a, b = b, a
	}

	x, y := 0, 0
	for i, j := len(a)-1, len(b)-1; i >= 0; {
		t := a[i] + x
		if j >= 0 {
			t += b[j]
		}

		a[i] = t % 2

		switch t {
		case 0, 1:
			x, y = y, 0
		default:
			x, y = y+1, 1
		}

		if x == 2 {
			x, y = 0, y-1
		}

		i--
		j--
	}

	if y != 0 {
		a = append([]int{y, x}, a...)
		return a
	} else if x != 0 {
		a = append([]int{x}, a...)
		return a
	}

	i := 0
	for i < len(a)-1 && a[i] == 0 {
		i++
	}
	a = a[i:]

	return a
}
