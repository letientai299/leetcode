package main

/*
 * @lc app=leetcode id=728 lang=golang
 *
 * [728] Self Dividing Numbers
 *
 * https://leetcode.com/problems/self-dividing-numbers/description/
 *
 * algorithms
 * Easy (70.69%)
 * Total Accepted:    86.1K
 * Total Submissions: 121.4K
 * Testcase Example:  '1\n22'
 *
 *
 * A self-dividing number is a number that is divisible by every digit it
 * contains.
 *
 * For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 ==
 * 0, and 128 % 8 == 0.
 *
 * Also, a self-dividing number is not allowed to contain the digit zero.
 *
 * Given a lower and upper number bound, output a list of every possible self
 * dividing number, including the bounds if possible.
 *
 * Example 1:
 *
 * Input:
 * left = 1, right = 22
 * Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
 *
 *
 *
 * Note:
 * The boundaries of each input argument are 1 .
 *
 */
func selfDividingNumbers(left int, right int) []int {
	if right < left {
		return nil
	}

	res := make([]int, 0, right-left)

	var f func(int) bool
	f = func(x int) bool {
		//m := make(map[int]bool)
		v := x
		for v > 0 {
			c := v % 10
			if c == 0 {
				return false
			}

			if x%c != 0 {
				return false
			}

			v = v / 10
		}
		return true
	}

	for v := left; v <= right; v++ {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
