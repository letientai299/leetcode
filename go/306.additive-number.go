package main

/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 *
 * https://leetcode.com/problems/additive-number/description/
 *
 * algorithms
 * Medium (28.82%)
 * Total Accepted:    55.5K
 * Total Submissions: 188.1K
 * Testcase Example:  '"112358"'
 *
 * Additive number is a string whose digits can form additive sequence.
 *
 * A valid additive sequence should contain at least three numbers. Except for
 * the first two numbers, each subsequent number in the sequence must be the
 * sum of the preceding two.
 *
 * Given a string containing only digits '0'-'9', write a function to determine
 * if it's an additive number.
 *
 * Note: Numbers in the additive sequence cannot have leading zeros, so
 * sequence 1, 2, 03 or 1, 02, 3 is invalid.
 *
 *
 * Example 1:
 *
 *
 * Input: "112358"
 * Output: true
 * Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5,
 * 8.
 * 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
 *
 *
 * Example 2:
 *
 *
 * Input: "199100199"
 * Output: true
 * Explanation: The additive sequence is: 1, 99, 100, 199.
 * 1 + 99 = 100, 99 + 100 = 199
 *
 *
 *
 * Constraints:
 *
 *
 * num consists only of digits '0'-'9'.
 * 1 <= num.length <= 35
 *
 *
 * Follow up:
 * How would you handle overflow for very large input integers?
 *
 */
func isAdditiveNumber(num string) bool {
	var isAdditive func(a, b int, s string) bool
	isAdditive = func(a, b int, s string) bool {
		if len(s) == 0 || (s[0] == '0' && a+b != 0) {
			return false
		}

		v := int(s[0] - '0')
		i := 1
		for ; i < len(s) && v < a+b; i++ {
			v = 10*v + int(s[i]-'0')
		}
		if v != a+b {
			return false
		}

		if i == len(s) {
			return true
		}

		return isAdditive(b, v, s[i:])
	}

	var valid func(arr []int, i int) bool

	valid = func(arr []int, i int) bool {
		c := len(arr)
		if i > (len(num)/3)*2 {
			return false
		}

		if c == 2 {
			if isAdditive(arr[0], arr[1], num[i:]) {
				return true
			}
		}

		d := int(num[i] - '0')
		if c == 0 {
			return valid(append(arr, d), i+1)
		}

		v := arr[c-1]
		if d == 0 {
			if v == 0 {
				if c < 2 {
					return valid(append(arr, 0), i+1)
				}
				return false
			}
		}

		if v == 0 {
			if c == 2 {
				return false
			}
		} else {
			v = 10*v + d
			arr[c-1] = v
			if valid(arr, i+1) {
				return true
			}
		}

		arr[c-1] = v / 10
		if c < 2 {
			arr := append(arr, d)
			res := valid(arr, i+1)
			return res
		}

		return false
	}

	if len(num) < 3 {
		return false
	}

	return valid([]int{}, 0)
}
