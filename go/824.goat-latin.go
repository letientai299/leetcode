package main

import "bytes"

/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 *
 * https://leetcode.com/problems/goat-latin/description/
 *
 * algorithms
 * Easy (59.77%)
 * Total Accepted:    103.4K
 * Total Submissions: 156.3K
 * Testcase Example:  '"I speak Goat Latin"'
 *
 * A sentence S is given, composed of words separated by spaces. Each word
 * consists of lowercase and uppercase letters only.
 *
 * We would like to convert the sentence to "Goat Latin" (a made-up language
 * similar to Pig Latin.)
 *
 * The rules of Goat Latin are as follows:
 *
 *
 * If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of
 * the word.
 * For example, the word 'apple' becomes 'applema'.
 *
 * If a word begins with a consonant (i.e. not a vowel), remove the first
 * letter and append it to the end, then add "ma".
 * For example, the word "goat" becomes "oatgma".
 *
 * Add one letter 'a' to the end of each word per its word index in the
 * sentence, starting with 1.
 * For example, the first word gets "a" added to the end, the second word gets
 * "aa" added to the end and so on.
 *
 *
 * Return the final sentence representing the conversion from S to Goat
 * Latin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "I speak Goat Latin"
 * Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
 *
 *
 * Example 2:
 *
 *
 * Input: "The quick brown fox jumped over the lazy dog"
 * Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa
 * hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
 *
 *
 *
 *
 * Notes:
 *
 *
 * S contains only uppercase, lowercase and spaces. Exactly one space between
 * each word.
 * 1 <= S.length <= 150.
 *
 *
 */
func toGoatLatin(s string) string {
	var buf bytes.Buffer
	prevIsLetter := false
	vowelFirst := false
	var consonant byte
	wordIndex := 0
	for i := range s {
		c := s[i]
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			if !prevIsLetter {
				vowelFirst = (c == 'a') || (c == 'e') || (c == 'i') || (c == 'o') || (c == 'u') ||
					(c == 'A') || (c == 'E') || (c == 'I') || (c == 'O') || (c == 'U')
				if !vowelFirst {
					consonant = c // skip this char, write later
				} else {
					buf.WriteByte(c)
				}
			} else {
				buf.WriteByte(c)
			}
			prevIsLetter = true
		} else {
			if !vowelFirst {
				buf.WriteByte(consonant)
			}
			buf.WriteString("ma")
			wordIndex++
			for x := 0; x < wordIndex; x++ {
				buf.WriteByte('a')
			}
			buf.WriteByte(c)
			prevIsLetter = false
		}
	}

	if prevIsLetter {
		if !vowelFirst {
			buf.WriteByte(consonant)
		}
		buf.WriteString("ma")
		wordIndex++
		for x := 0; x < wordIndex; x++ {
			buf.WriteByte('a')
		}
	}

	return buf.String()
}
