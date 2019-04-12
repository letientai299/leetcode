package main

/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 *
 * https://leetcode.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (61.90%)
 * Total Accepted:    85.3K
 * Total Submissions: 137.8K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * Given a List of words, return the words that can be typed using letters of
 * alphabet on only one row's of American keyboard like the image
 * below.
 *
 *
 *
 *
 *
 *
 * Example:
 *
 *
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 *
 *
 *
 *
 * Note:
 *
 *
 * You may use one character in the keyboard more than once.
 * You may assume the input string will only contain letters of alphabet.
 *
 *
 */
func findWords(words []string) []string {
	rows := map[int32]int{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,

		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,

		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
	}

	var res []string
	for _, w := range words {
		u := w[0]
		if u <= 'Z' {
			u += 'a' - 'A'
		}
		r := rows[int32(u)]
		ok := true
		for _, c := range w {
			if c <= 'Z' {
				c += 'a' - 'A'
			}

			if rows[c] != r {
				ok = false
				break
			}
		}

		if ok {
			res = append(res, w)
		}
	}
	return res
}
