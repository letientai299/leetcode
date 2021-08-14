package main

import "strings"

// Print Words Vertically
//
// Medium
//
// https://leetcode.com/problems/print-words-vertically/
//
// Given a string `s`. Return all the words vertically in the same order in
// which they appear in `s`.
//
// Words are returned as a list of strings, complete with spaces when is
// necessary. (Trailing spaces are not allowed).
//
// Each word would be put on only one column and that in one column there will
// be only one word.
//
// **Example 1:**
//
// ```
// Input: s = "HOW ARE YOU"
// Output: ["HAY","ORO","WEU"]
// Explanation: Each word is printed vertically.
//  "HAY"
//  "ORO"
//  "WEU"
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "TO BE OR NOT TO BE"
// Output: ["TBONTB","OEROOE","   T"]
// Explanation: Trailing spaces is not allowed.
// "TBONTB"
// "OEROOE"
// "   T"
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "CONTEST IS COMING"
// Output: ["CIC","OSO","N M","T I","E N","S G","T"]
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 200`
// - `s` contains only upper case English letters.
// - It's guaranteed that there is only one space between 2 words.
func printVertically(s string) []string {
	var res []string
	ss := strings.Split(s, " ")
	i := 0
	ok := true
	var sb strings.Builder
	for ok {
		ok = false
		for _, w := range ss {
			if i < len(w) {
				ok = true
				sb.WriteByte(w[i])
			} else {
				sb.WriteByte(' ')
			}
		}
		if ok {
			res = append(res, strings.TrimRight(sb.String(), " "))
			sb.Reset()
			i++
		}
	}

	return res
}