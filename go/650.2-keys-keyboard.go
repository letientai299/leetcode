package main

import (
	"math"
)

/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 *
 * https://leetcode.com/problems/2-keys-keyboard/description/
 *
 * algorithms
 * Medium (47.51%)
 * Total Accepted:    66K
 * Total Submissions: 132.9K
 * Testcase Example:  '3'
 *
 * Initially on a notepad only one character 'A' is present. You can perform
 * two operations on this notepad for each step:
 *
 *
 * Copy All: You can copy all the characters present on the notepad (partial
 * copy is not allowed).
 * Paste: You can paste the characters which are copied last time.
 *
 *
 *
 *
 * Given a number n. You have to get exactly n 'A' on the notepad by performing
 * the minimum number of steps permitted. Output the minimum number of steps to
 * get n 'A'.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation:
 * Intitally, we have one character 'A'.
 * In step 1, we use Copy All operation.
 * In step 2, we use Paste operation to get 'AA'.
 * In step 3, we use Paste operation to get 'AAA'.
 *
 *
 *
 *
 * Note:
 *
 *
 * The n will be in the range [1, 1000].
 *
 *
 *
 *
 */

func minSteps650(n int) int {
	if n < 2 {
		return 0
	}
	if n >= 2 && n <= 5 {
		return n
	}

	r := 0
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		for n%i == 0 {
			r += i
			n /= i
		}
	}

	if n > 1 {
		r += n
	}

	return r

	// f(1) = 0
	// f(2) = 2
	// f(3) = 3
	// f(4) = 4
	// f(5) = 5
	// f(6) = 5
	// f(7) = 7
	// f(7) = 7
	// f(8) = 6
	// f(9) = 6
	// f(10) = 7
	// f(11) = 11
	//
	// f(xn) = f(n)+x
	// f(2^x n) = f(n)+2x
}
