package main

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 *
 * https://leetcode.com/problems/string-compression/description/
 *
 * algorithms
 * Easy (37.21%)
 * Total Accepted:    46.6K
 * Total Submissions: 125.3K
 * Testcase Example:  '["a","a","b","b","c","c","c"]'
 *
 * Given an array of characters, compress it in-place.
 *
 * The length after compression must always be smaller than or equal to the
 * original array.
 *
 * Every element of the array should be a character (not int) of length 1.
 *
 * After you are done modifying the input array in-place, return the new length
 * of the array.
 *
 *
 * Follow up:
 * Could you solve it using only O(1) extra space?
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["a","a","b","b","c","c","c"]
 *
 * Output:
 * Return 6, and the first 6 characters of the input array should be:
 * ["a","2","b","2","c","3"]
 *
 * Explanation:
 * "aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by
 * "c3".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ["a"]
 *
 * Output:
 * Return 1, and the first 1 characters of the input array should be: ["a"]
 *
 * Explanation:
 * Nothing is replaced.
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input:
 * ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
 *
 * Output:
 * Return 4, and the first 4 characters of the input array should be:
 * ["a","b","1","2"].
 *
 * Explanation:
 * Since the character "a" does not repeat, it is not compressed.
 * "bbbbbbbbbbbb" is replaced by "b12".
 * Notice each digit has it's own entry in the array.
 *
 *
 *
 *
 * Note:
 *
 *
 * All characters have an ASCII value in [35, 126].
 * 1 <= len(chars) <= 1000.
 *
 *
 */
func compress(s []byte) int {
	if len(s) < 2 {
		return len(s)
	}

	current := 1
	location := 0
	f := func(i int) {
		s[location] = s[i]
		switch current {
		case 1:
			location++
		case 2:
			s[location+1] = '2'
			location += 2
		default:
			location++
			upper := 1
			for upper <= current {
				upper *= 10
			}

			upper /= 10

			for current >= 10 {
				x := current / upper
				current %= upper
				upper /= 10
				s[location] = byte(x + '0')
				location++
			}
			s[location] = byte(current + '0')
			location++
		}
		current = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			current++
			continue
		}

		f(i - 1)
	}
	f(len(s) - 1)
	return location
}
