package main

/*
 * @lc app=leetcode id=422 lang=golang
 *
 * [422] Valid Word Square
 *
 * https://leetcode.com/problems/valid-word-square/description/
 *
 * algorithms
 * Easy (36.83%)
 * Total Accepted:    31.1K
 * Total Submissions: 81.9K
 * Testcase Example:  '["abcd","bnrt","crmy","dtye"]'
 *
 * Given a sequence of words, check whether it forms a valid word square.
 *
 * A sequence of words forms a valid word square if the k^th row and column
 * read the exact same string, where 0 ≤ k < max(numRows, numColumns).
 *
 * Note:
 *
 * The number of words given is at least 1 and does not exceed 500.
 * Word length will be at least 1 and does not exceed 500.
 * Each word contains only lowercase English alphabet a-z.
 *
 *
 *
 * Example 1:
 *
 * Input:
 * [
 * ⁠ "abcd",
 * ⁠ "bnrt",
 * ⁠ "crmy",
 * ⁠ "dtye"
 * ]
 *
 * Output:
 * true
 *
 * Explanation:
 * The first row and first column both read "abcd".
 * The second row and second column both read "bnrt".
 * The third row and third column both read "crmy".
 * The fourth row and fourth column both read "dtye".
 *
 * Therefore, it is a valid word square.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ "abcd",
 * ⁠ "bnrt",
 * ⁠ "crm",
 * ⁠ "dt"
 * ]
 *
 * Output:
 * true
 *
 * Explanation:
 * The first row and first column both read "abcd".
 * The second row and second column both read "bnrt".
 * The third row and third column both read "crm".
 * The fourth row and fourth column both read "dt".
 *
 * Therefore, it is a valid word square.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * [
 * ⁠ "ball",
 * ⁠ "area",
 * ⁠ "read",
 * ⁠ "lady"
 * ]
 *
 * Output:
 * false
 *
 * Explanation:
 * The third row reads "read" while the third column reads "lead".
 *
 * Therefore, it is NOT a valid word square.
 *
 *
 */
func validWordSquare(words []string) bool {
	if len(words) == 0 {
		return true
	}

	for x := 0; x < len(words); x++ {
		for y := 0; y < len(words[x]); y++ {
			if y >= len(words) || x >= len(words[y]) {
				return false
			}

			if words[x][y] != words[y][x] {
				return false
			}
		}
	}

	return true
}
