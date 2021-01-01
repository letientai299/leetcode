package main

/*
 * @lc app=leetcode id=481 lang=golang
 *
 * [481] Magical String
 *
 * https://leetcode.com/problems/magical-string/description/
 *
 * algorithms
 * Medium (46.72%)
 * Total Accepted:    23.5K
 * Total Submissions: 49.1K
 * Testcase Example:  '1'
 *
 *
 * A magical string S consists of only '1' and '2' and obeys the following
 * rules:
 *
 *
 * The string S is magical because concatenating the number of contiguous
 * occurrences of characters '1' and '2' generates the string S itself.
 *
 *
 *
 * The first few elements of string S is the following:
 * S = "1221121221221121122……"
 *
 *
 *
 * If we group the consecutive '1's and '2's in S, it will be:
 *
 *
 * 1   22  11  2  1  22  1  22  11  2  11  22 ......
 *
 *
 * and the occurrences of '1's or '2's in each group are:
 *
 *
 * 1   2       2    1   1    2     1    2     2    1    2    2 ......
 *
 *
 *
 * You can see that the occurrence sequence above is the S itself.
 *
 *
 *
 * Given an integer N as input, return the number of '1's in the first N number
 * in the magical string S.
 *
 *
 * Note:
 * N will not exceed 100,000.
 *
 *
 *
 * Example 1:
 *
 * Input: 6
 * Output: 3
 * Explanation: The first 6 elements of magical string S is "12211" and it
 * contains three 1's, so return 3.
 *
 *
 */
func magicalString(n int) int {
	if n < 4 {
		if n == 0 {
			return 0
		}
		return 1
	}

	buf := make([]byte, 0, n)
	buf = append(buf, 1, 2, 2)
	r := 1
	for i := 2; len(buf) < n; i++ {
		c := 3 - buf[len(buf)-1]
		if buf[i] == 1 {
			buf = append(buf, c)
		} else {
			buf = append(buf, c, c)
		}
		if c == 1 {
			r += int(buf[i])
			if len(buf) > n {
				r -= len(buf) - n
			}
		}
	}
	return r
}
