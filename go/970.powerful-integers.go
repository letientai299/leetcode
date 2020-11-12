package main

/*
 * @lc app=leetcode id=970 lang=golang
 *
 * [970] Powerful Integers
 *
 * https://leetcode.com/problems/powerful-integers/description/
 *
 * algorithms
 * Easy (39.65%)
 * Total Accepted:    24.6K
 * Total Submissions: 61.6K
 * Testcase Example:  '2\n3\n10'
 *
 * Given two positive integers x and y, an integer is powerful if it is equal
 * to x^i + y^j for some integers i >= 0 and j >= 0.
 *
 * Return a list of all powerful integers that have value less than or equal to
 * bound.
 *
 * You may return the answer in any order.  In your answer, each value should
 * occur at most once.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: x = 2, y = 3, bound = 10
 * Output: [2,3,4,5,7,9,10]
 * Explanation:
 * 2 = 2^0 + 3^0
 * 3 = 2^1 + 3^0
 * 4 = 2^0 + 3^1
 * 5 = 2^1 + 3^1
 * 7 = 2^2 + 3^1
 * 9 = 2^3 + 3^0
 * 10 = 2^0 + 3^2
 *
 *
 *
 * Example 2:
 *
 *
 * Input: x = 3, y = 5, bound = 15
 * Output: [2,4,6,8,10,14]
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= x <= 100
 * 1 <= y <= 100
 * 0 <= bound <= 10^6
 *
 */
func powerfulIntegers(x int, y int, bound int) []int {
	xs := []int{1}
	ys := []int{1}
	rx, ry := x, y

	for (rx < bound && rx != 1) || (ry != 1 && ry < bound) {
		if rx < bound && rx != 1 {
			xs = append(xs, rx)
			rx *= x
		}

		if ry < bound && ry != 1 {
			ys = append(ys, ry)
			ry *= y
		}
	}

	var res []int
	m := make(map[int]struct{})

	for _, a := range xs {
		for _, b := range ys {
			v := a + b
			if v > bound {
				continue
			}

			if _, ok := m[v]; !ok {
				res = append(res, v)
				m[v] = struct{}{}
			}
		}
	}

	return res
}
