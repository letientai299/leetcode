package main

import "math/bits"

/*
 * @lc app=leetcode id=762 lang=golang
 *
 * [762] Prime Number of Set Bits in Binary Representation
 *
 * https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/
 *
 * algorithms
 * Easy (60.66%)
 * Total Accepted:    34K
 * Total Submissions: 56K
 * Testcase Example:  '842\n888'
 *
 *
 * Given two integers L and R, find the count of numbers in the range [L, R]
 * (inclusive) having a prime number of set bits in their binary
 * representation.
 *
 * (Recall that the number of set bits an integer has is the number of 1s
 * present when written in binary.  For example, 21 written in binary is 10101
 * which has 3 set bits.  Also, 1 is not a prime.)
 *
 *
 * Example 1:
 * Input: L = 6, R = 10
 * Output: 4
 * Explanation:
 * 6 -> 110 (2 set bits, 2 is prime)
 * 7 -> 111 (3 set bits, 3 is prime)
 * 9 -> 1001 (2 set bits , 2 is prime)
 * 10->1010 (2 set bits , 2 is prime)
 *
 *
 * Example 2:
 * Input: L = 10, R = 15
 * Output: 5
 * Explanation:
 * 10 -> 1010 (2 set bits, 2 is prime)
 * 11 -> 1011 (3 set bits, 3 is prime)
 * 12 -> 1100 (2 set bits, 2 is prime)
 * 13 -> 1101 (3 set bits, 3 is prime)
 * 14 -> 1110 (3 set bits, 3 is prime)
 * 15 -> 1111 (4 set bits, 4 is not prime)
 *
 *
 * Note:
 * L, R will be integers L  in the range [1, 10^6].
 * R - L will be at most 10000.
 *
 */

func countPrimeSetBits(L int, R int) int {
	// we know that there's maximum 32 bits
	// thus, there's a limited possible of primes
	primes := map[int]struct{}{
		2:  {},
		3:  {},
		5:  {},
		7:  {},
		11: {},
		13: {},
		17: {},
		19: {},
		23: {},
		29: {},
		31: {},
	}

	count := 0
	for x := L; x <= R; x++ {
		b := bits.OnesCount(uint(x))
		if _, ok := primes[b]; ok {
			count++
		}
	}
	return count
}

// TODO (jl): below algorithm should be faster than the loop through L to R
//     but I'm too lazy
/*
func countPrimeSetBits(L int, R int) int {
	// find number of bits of L and R
	lb := 0
	for L > 0 {
		lb++
		L >>= 1
	}

	rb := 0
	for R > 0 {
		rb++
		R >>= 1
	}

	// we know that there's maximum 32 bits
	// thus, there's a limited possible of primes
	primes := map[int]struct{}{
		2:  {},
		3:  {},
		5:  {},
		7:  {},
		11: {},
		13: {},
		19: {},
		23: {},
		29: {},
		31: {},
	}

	count := 0
	for i := lb; i <= rb; i++ {
		if _, ok := primes[i]; ok {
			// for a set of bits, we can construct all numbers that has
			// at least lb bits and at most rb bits.
			// then, if that number is between L and R, we increase the counter.
			count +=
		}
	}

	return count
}
*/
