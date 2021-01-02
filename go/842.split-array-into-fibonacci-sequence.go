package main

import (
	"math"
)

/*
 * @lc app=leetcode id=842 lang=golang
 *
 * [842] Split Array into Fibonacci Sequence
 *
 * https://leetcode.com/problems/split-array-into-fibonacci-sequence/description/
 *
 * algorithms
 * Medium (35.45%)
 * Total Accepted:    23K
 * Total Submissions: 62.9K
 * Testcase Example:  '"123456579"'
 *
 * Given a string S of digits, such as S = "123456579", we can split it into a
 * Fibonacci-like sequence [123, 456, 579].
 *
 * Formally, a Fibonacci-like sequence is a list F of non-negative integers
 * such that:
 *
 *
 * 0 <= F[i] <= 2^31 - 1, (that is, each integer fits a 32-bit signed integer
 * type);
 * F.length >= 3;
 * and F[i] + F[i+1] = F[i+2] for all 0 <= i < F.length - 2.
 *
 *
 * Also, note that when splitting the string into pieces, each piece must not
 * have extra leading zeroes, except if the piece is the number 0 itself.
 *
 * Return any Fibonacci-like sequence split from S, or return [] if it cannot
 * be done.
 *
 * Example 1:
 *
 *
 * Input: "123456579"
 * Output: [123,456,579]
 *
 *
 * Example 2:
 *
 *
 * Input: "11235813"
 * Output: [1,1,2,3,5,8,13]
 *
 *
 * Example 3:
 *
 *
 * Input: "112358130"
 * Output: []
 * Explanation: The task is impossible.
 *
 *
 * Example 4:
 *
 *
 * Input: "0123"
 * Output: []
 * Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not
 * valid.
 *
 *
 * Example 5:
 *
 *
 * Input: "1101111"
 * Output: [110, 1, 111]
 * Explanation: The output [11, 0, 11, 11] would also be accepted.
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 200
 * S contains only digits.
 *
 *
 */
func splitIntoFibonacci(num string) []int {
	if len(num) < 3 {
		return nil
	}

	var checkFibo func(s string, cur []int) []int
	checkFibo = func(s string, cur []int) []int {
		n := len(cur)
		a, b := cur[n-2], cur[n-1]
		if len(s) == 0 || (s[0] == '0' && a+b != 0) {
			return nil
		}

		v := int(s[0] - '0')
		i := 1
		for ; i < len(s) && v < a+b; i++ {
			v = 10*v + int(s[i]-'0')
		}

		if v != a+b {
			return nil
		}

		if v > math.MaxInt32 {
			return nil
		}

		cur = append(cur, v)
		if i == len(s) {
			return cur
		}

		return checkFibo(s[i:], cur)
	}

	var splitFibo func(arr []int, i int) []int

	splitFibo = func(arr []int, i int) []int {
		c := len(arr)
		if i >= len(num) {
			return nil
		}

		if c == 2 {
			res := checkFibo(num[i:], arr)
			if len(res) != 0 {
				return res
			}
		}

		d := int(num[i] - '0')
		if c == 0 {
			return splitFibo(append(arr, d), i+1)
		}

		v := arr[c-1]
		if d == 0 {
			if v == 0 {
				if c < 2 {
					return splitFibo(append(arr, 0), i+1)
				}
				return nil
			}
		}

		if v == 0 {
			if c == 2 {
				return nil
			}
		} else {
			v = 10*v + d
			if v <= math.MaxInt32 {
				arr[c-1] = v
				res := splitFibo(arr, i+1)
				if len(res) != 0 {
					return res
				}
			}
		}

		arr[c-1] = v / 10
		if c < 2 {
			arr := append(arr, d)
			return splitFibo(arr, i+1)
		}

		return nil
	}

	return splitFibo([]int{}, 0)
}
