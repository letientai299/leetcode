package main

/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 *
 * https://leetcode.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (59.61%)
 * Total Accepted:    60.4K
 * Total Submissions: 101.4K
 * Testcase Example:  '"a1b2"'
 *
 * Given a string S, we can transform every letter individually to be lowercase
 * or uppercase to create another string.  Return a list of all possible
 * strings we could create.
 *
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 *
 * Note:
 *
 *
 * S will be a string with length between 1 and 12.
 * S will consist only of letters or digits.
 *
 *
 */
func letterCasePermutation(S string) []string {
	bf := make([]byte, 0, len(S))
	for i := range S {
		c := S[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'a' - 'A'
		}

		bf = append(bf, c)
	}

	// use []byte for faster perf
	all := make([][]byte, 0, 1<<uint32(len(bf)))
	all = append(all, bf)
	for i, c := range bf {
		if c < 'a' || c > 'z' {
			continue
		}

		news := make([][]byte, 0, len(all))
		for _, v := range all {
			b := make([]byte, len(bf))
			copy(b, v)
			b[i] = c - 'a' + 'A'
			news = append(news, b)
		}
		all = append(all, news...)
	}

	// convert byte to string
	ss := make([]string, 0, len(all))
	for _, x := range all {
		ss = append(ss, string(x))
	}

	return ss
}
