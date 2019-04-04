package main

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 *
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (36.81%)
 * Total Accepted:    114.3K
 * Total Submissions: 310.5K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * Given a string s and a non-empty string p, find all the start indices of p's
 * anagrams in s.
 *
 * Strings consists of lowercase English letters only and the length of both
 * strings s and p will not be larger than 20,100.
 *
 * The order of output does not matter.
 *
 * Example 1:
 *
 * Input:
 * s: "cbaebabacd" p: "abc"
 *
 * Output:
 * [0, 6]
 *
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of
 * "abc".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s: "abab" p: "ab"
 *
 * Output:
 * [0, 1, 2]
 *
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 *
 *
 */
func findAnagrams(s string, p string) []int {
	n := len(p)
	if len(s) < n {
		return nil
	}

	pc := make([]int, 26)
	sc := make([]int, 26)
	for i, c := range p {
		pc[c-'a']++
		sc[s[i]-'a']++
	}

	var res []int
	found := true
	for x := 0; x < len(pc); x++ {
		if pc[x] != sc[x] {
			found = false
			break
		}
	}

	if found {
		res = append(res, 0)
	}

	for i := n; i < len(s); i++ {
		sc[s[i-n]-'a']--
		sc[s[i]-'a']++

		found := true
		for x := 0; x < len(pc); x++ {
			if pc[x] != sc[x] {
				found = false
				break
			}
		}

		if found {
			res = append(res, i-n+1)
		}
	}

	return res
}
