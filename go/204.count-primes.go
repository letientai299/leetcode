package main

import "sort"

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
var knowns = []int{2}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	max := sort.Search(len(knowns), func(i int) bool { return knowns[i] >= n })
	if max < len(knowns) {
		return max
	}

	for i := knowns[len(knowns)-1] + 1; i < n; i++ {
		isPrime := true
		for _, x := range knowns {
			if x*x > i {
				break
			}

			if i%x == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			knowns = append(knowns, i)
		}
	}

	return len(knowns)
}
