package main

/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 *
 * https://leetcode.com/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (37.33%)
 * Total Accepted:    196.8K
 * Total Submissions: 478.9K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G',
 * and 'T', for example: "ACGAATTCCG". When studying DNA, it is sometimes
 * useful to identify repeated sequences within the DNA.
 *
 * Write a function to find all the 10-letter-long sequences (substrings) that
 * occur more than once in a DNA molecule.
 *
 *
 * Example 1:
 * Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 * Output: ["AAAAACCCCC","CCCCCAAAAA"]
 * Example 2:
 * Input: s = "AAAAAAAAAAAAA"
 * Output: ["AAAAAAAAAA"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 10^5
 * s[i] is 'A', 'C', 'G', or 'T'.
 *
 *
 */
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return nil
	}

	var digits ['T' - 'A' + 1]int64
	digits['A'-'A'] = 1
	digits['C'-'A'] = 2
	digits['G'-'A'] = 3
	digits['T'-'A'] = 4

	w := int64(0)
	i := 0
	for ; i < 10; i++ {
		w = 10*w + digits[s[i]-'A']
	}

	count := make(map[int64]int)
	count[w] = 1

	var res []string
	for ; i < n; i++ {
		w = (w-1e9*digits[s[i-10]-'A'])*10 + digits[s[i]-'A']
		count[w]++
		if count[w] == 2 {
			res = append(res, s[i-9:i+1])
		}
	}

	return res
}
