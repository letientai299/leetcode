package main

/*
 * @lc app=leetcode id=809 lang=golang
 *
 * [809] Expressive Words
 *
 * https://leetcode.com/problems/expressive-words/description/
 *
 * algorithms
 * Medium (45.58%)
 * Total Accepted:    59.1K
 * Total Submissions: 126.8K
 * Testcase Example:  '"heeellooo"\n["hello", "hi", "helo"]'
 *
 * Sometimes people repeat letters to represent extra feeling, such as "hello"
 * -> "heeellooo", "hi" -> "hiiii".  In these strings like "heeellooo", we have
 * groups of adjacent letters that are all the same:  "h", "eee", "ll", "ooo".
 *
 * For some given string S, a query word is stretchy if it can be made to be
 * equal to S by any number of applications of the following extension
 * operation: choose a group consisting of characters c, and add some number of
 * characters c to the group so that the size of the group is 3 or more.
 *
 * For example, starting with "hello", we could do an extension on the group
 * "o" to get "hellooo", but we cannot get "helloo" since the group "oo" has
 * size less than 3.  Also, we could do another extension like "ll" -> "lllll"
 * to get "helllllooo".  If S = "helllllooo", then the query word "hello" would
 * be stretchy because of these two extension operations: query = "hello" ->
 * "hellooo" -> "helllllooo" = S.
 *
 * Given a list of query words, return the number of words that are
 * stretchy. 
 *
 *
 *
 *
 * Example:
 * Input:
 * S = "heeellooo"
 * words = ["hello", "hi", "helo"]
 * Output: 1
 * Explanation:
 * We can extend "e" and "o" in the word "hello" to get "heeellooo".
 * We can't extend "helo" to get "heeellooo" because the group "ll" is not size
 * 3 or more.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= len(S) <= 100.
 * 0 <= len(words) <= 100.
 * 0 <= len(words[i]) <= 100.
 * S and all words in words consist only of lowercase letters
 *
 *
 */

func expressiveWords(s string, words []string) int {
	if len(s) == 0 {
		return 0
	}

	count := func(s string) []int {
		m := make([]int, 0, len(s)/2)
		c := s[0]
		m = append(m, int(c))
		m = append(m, 1)
		for i := 1; i < len(s); i++ {
			if c == s[i] {
				m[len(m)-1]++
			} else {
				c = s[i]
				m = append(m, int(s[i]))
				m = append(m, 1)
			}
		}
		return m
	}

	base := count(s)

	res := 0
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		c := w[0]
		j := 0
		if c != s[j] {
			continue
		}

		n := 1
		ok := true
		for i := 1; i < len(w) && ok && j < len(base)-1; i++ {
			if c == w[i] {
				n++
				ok = n <= base[j+1]
				continue
			}

			if base[j+1] < 3 && base[j+1] != n {
				ok = false
				break
			}

			c = w[i]
			n = 1
			j += 2
		}

		if !ok || j != len(base)-2 || base[j+1] < n || (base[j+1] == 2 && n == 1) {
			continue
		}

		res++
	}
	return res
}