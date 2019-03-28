package main

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (41.02%)
 * Total Accepted:    146.7K
 * Total Submissions: 357.7K
 * Testcase Example:  '"hello"'
 *
 * Write a function that takes a string as input and reverse only the vowels of
 * a string.
 *
 * Example 1:
 *
 *
 * Input: "hello"
 * Output: "holle"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "leetcode"
 * Output: "leotcede"
 *
 *
 * Note:
 * The vowels does not include the letter "y".
 *
 *
 *
 */
func reverseVowels(str string) string {
	s := []byte(str)
	isVowel := func(c uint8) bool {
		switch c {
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case 'e':
			fallthrough
		case 'E':
			fallthrough
		case 'i':
			fallthrough
		case 'I':
			fallthrough
		case 'o':
			fallthrough
		case 'O':
			fallthrough
		case 'u':
			fallthrough
		case 'U':
			return true
		default:
			return false
		}
	}

	x := 0
	n := len(s)
	y := n - 1
	for x < y {
		if !isVowel(s[x]) {
			x++
		}

		if !isVowel(s[y]) {
			y--
		}

		if x < y && isVowel(s[x]) && isVowel(s[y]) {
			s[x], s[y] = s[y], s[x]
			x++
			y--
		}
	}

	return string(s)
}
