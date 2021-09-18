package main

/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (28.46%)
 * Total Accepted:    220.1K
 * Total Submissions: 773.6K
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Example:
 *
 *
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 *
 */

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	if n <= 3 {
		return 1
	}

	notPrimes := make([]bool, n)
	notPrimes[2] = false

	r := 1
	for i := 3; i < n; i += 2 {
		if !notPrimes[i] {
			r++
			// 3 9 15
			for j := 3 * i; j < n; j += 2 * i {
				notPrimes[j] = true
			}
		}
	}

	return r
}
