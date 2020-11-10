package main

/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 *
 * https://leetcode.com/problems/verifying-an-alien-dictionary/description/
 *
 * algorithms
 * Easy (55.36%)
 * Total Accepted:    159.1K
 * Total Submissions: 300.1K
 * Testcase Example:  '["hello","leetcode"]\n"hlabcdefgijkmnopqrstuvwxyz"'
 *
 * In an alien language, surprisingly they also use english lowercase letters,
 * but possibly in a different order. The order of the alphabet is some
 * permutation of lowercase letters.
 *
 * Given a sequence of words written in the alien language, and the order of
 * the alphabet, return true if and only if the given words are sorted
 * lexicographicaly in this alien language.
 *
 * Example 1:
 *
 *
 * Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
 * Output: true
 * Explanation: As 'h' comes before 'l' in this language, then the sequence is
 * sorted.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
 * Output: false
 * Explanation: As 'd' comes after 'l' in this language, then words[0] >
 * words[1], hence the sequence is unsorted.
 *
 *
 * Example 3:
 *
 *
 * Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
 * Output: false
 * Explanation: The first three characters "app" match, and the second string
 * is shorter (in size.) According to lexicographical rules "apple" > "app",
 * because 'l' > '∅', where '∅' is defined as the blank character which is less
 * than any other character (More info).
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 100
 * 1 <= words[i].length <= 20
 * order.length == 26
 * All characters in words[i] and order are English lowercase letters.
 *
 *
 */
func isAlienSorted(words []string, order string) bool {
	rank := [26]int{}
	for i, c := range order {
		rank[c-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		a, b := words[i-1], words[i]
		less := false
		for x := 0; x < len(a) && x < len(b); x++ {
			if rank[a[x]-'a'] < rank[b[x]-'a'] {
				less = true
				break
			}

			if rank[a[x]-'a'] == rank[b[x]-'a'] {
				continue
			}

			if rank[a[x]-'a'] > rank[b[x]-'a'] {
				return false
			}
		}

		if !less && len(a) > len(b) {
			return false
		}
	}
	return true
}
