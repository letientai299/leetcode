package main

import "bytes"

/*
* @lc app=leetcode id=819 lang=golang
*
* [819] Most Common Word
*
* https://leetcode.com/problems/most-common-word/description/
*
* algorithms
* Easy (43.17%)
* Total Accepted:    190.7K
* Total Submissions: 421.9K
* Testcase Example:  '"Bob hit a ball, the hit BALL flew far after it was hit."\n["hit"]'
*
* Given a paragraph and a list of banned words, return the most frequent word
* that is not in the list of banned words.  It is guaranteed there is at least
* one word that isn't banned, and that the answer is unique.
*
* Words in the list of banned words are given in lowercase, and free of
* punctuation.  Words in the paragraph are not case sensitive.  The answer is
* in lowercase.
*
*
*
* Example:
*
*
* Input:
* paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
* banned = ["hit"]
* Output: "ball"
* Explanation:
* "hit" occurs 3 times, but it is a banned word.
* "ball" occurs twice (and no other word does), so it is the most frequent
* non-banned word in the paragraph.
* Note that words in the paragraph are not case sensitive,
* that punctuation is ignored (even if adjacent to words, such as "ball,"),
* and that "hit" isn't the answer even though it occurs more because it is
* banned.
*
*
*
*
* Note:
*
*
* 1 <= paragraph.length <= 1000.
* 0 <= banned.length <= 100.
* 1 <= banned[i].length <= 10.
* The answer is unique, and written in lowercase (even if its occurrences in
* paragraph may have uppercase symbols, and even if it is a proper noun.)
* paragraph only consists of letters, spaces, or the punctuation symbols
* !?',;.
* There are no hyphens or hyphenated words.
* Words only consist of letters, never apostrophes or other punctuation
* symbols.
*
*
 */
func mostCommonWord(paragraph string, banned []string) string {
	ban := make(map[string]struct{}, len(banned))
	for _, w := range banned {
		ban[w] = struct{}{}
	}

	m := make(map[string]int)
	top := ""
	var buf bytes.Buffer
	for _, c := range paragraph {
		switch {
		case c >= 'a' && c <= 'z':
			buf.WriteByte(byte(c))
		case c >= 'A' && c <= 'Z':
			buf.WriteByte(byte(c - 'A' + 'a'))
		default:
			if buf.Len() == 0 {
				continue
			}
			s := buf.String()
			buf.Reset()
			if _, ok := ban[s]; ok {
				continue
			}
			m[s]++
			if m[top] < m[s] {
				top = s
			}
		}
	}

	if buf.Len() != 0 {
		s := buf.String()
		buf.Reset()
		if _, ok := ban[s]; !ok {
			m[s]++
			if m[top] < m[s] {
				top = s
			}
		}
	}

	return top
}
