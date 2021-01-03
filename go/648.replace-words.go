package main

import "strings"

/*
 * @lc app=leetcode id=648 lang=golang
 *
 * [648] Replace Words
 *
 * https://leetcode.com/problems/replace-words/description/
 *
 * algorithms
 * Medium (54.02%)
 * Total Accepted:    64.5K
 * Total Submissions: 111K
 * Testcase Example:  '["cat","bat","rat"]\n"the cattle was rattled by the battery"'
 *
 * In English, we have a concept called root, which can be followed by some
 * other word to form another longer word - let's call this word successor. For
 * example, when the root "an" is followed by the successor word "other", we
 * can form a new word "another".
 *
 * Given a dictionary consisting of many roots and a sentence consisting of
 * words separated by spaces, replace all the successors in the sentence with
 * the root forming it. If a successor can be replaced by more than one root,
 * replace it with the root that has the shortest length.
 *
 * Return the sentence after the replacement.
 *
 *
 * Example 1:
 * Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled
 * by the battery"
 * Output: "the cat was rat by the bat"
 * Example 2:
 * Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
 * Output: "a a b c"
 * Example 3:
 * Input: dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa
 * aaa aaa aaaaaa bbb baba ababa"
 * Output: "a a a a a a a a bbb baba a"
 * Example 4:
 * Input: dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was
 * rattled by the battery"
 * Output: "the cat was rat by the bat"
 * Example 5:
 * Input: dictionary = ["ac","ab"], sentence = "it is abnormal that this
 * solution is accepted"
 * Output: "it is ab that this solution is ac"
 *
 *
 * Constraints:
 *
 *
 * 1 <= dictionary.length <= 1000
 * 1 <= dictionary[i].length <= 100
 * dictionary[i] consists of only lower-case letters.
 * 1 <= sentence.length <= 10^6
 * sentence consists of only lower-case letters and spaces.
 * The number of words in sentence is in the range [1, 1000]
 * The length of each word in sentence is in the range [1, 1000]
 * Each two consecutive words in sentence will be separated by exactly one
 * space.
 * sentence does not have leading or trailing spaces.
 *
 *
 */
func replaceWords(dictionary []string, sentence string) string {
	t := newTrie()
	for _, w := range dictionary {
		t.insert(w)
	}

	var sb strings.Builder
	var buf []byte
	for _, c := range sentence {
		if c == ' ' {
			w := string(buf)
			buf = buf[:0]
			origin, ok := t.findPrefix(w)
			if !ok {
				sb.WriteString(w)
			} else {
				sb.WriteString(origin)
			}
			sb.WriteByte(' ')
			continue
		}

		buf = append(buf, byte(c))
	}

	w := string(buf)
	origin, ok := t.findPrefix(w)
	if !ok {
		sb.WriteString(w)
	} else {
		sb.WriteString(origin)
	}
	return sb.String()
}
