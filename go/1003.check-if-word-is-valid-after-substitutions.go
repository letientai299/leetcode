package main

// Check If Word Is Valid After Substitutions
//
// Medium
//
// https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
//
// Given a string `s`, determine if it is **valid**.
//
// A string `s` is **valid** if, starting with an empty string `t = ""`, you can
// **transform** `t` **into** `s` after performing the following operation **any
// number of times**:
//
// - Insert string `"abc"` into any position in `t`. More formally, `t` becomes
// `tleft + "abc" + tright`, where `t == tleft + tright`. Note that `tleft` and
// `tright` may be **empty**.
//
// Return `true` _if_ `s` _is a **valid** string, otherwise, return_ `false`.
//
// **Example 1:**
//
// ```
// Input: s = "aabcbc"
// Output: true
// Explanation:
// "" -> "abc" -> "aabcbc"
// Thus, "aabcbc" is valid.
// ```
//
// **Example 2:**
//
// ```
// Input: s = "abcabcababcc"
// Output: true
// Explanation:
// "" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
// Thus, "abcabcababcc" is valid.
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "abccba"
// Output: false
// Explanation: It is impossible to get "abccba" using the operation.
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = "cababc"
// Output: false
// Explanation: It is impossible to get "cababc" using the operation.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 2 * 104`
// - `s` consists of letters `'a'`, `'b'`, and `'c'`
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, c := range s {
		if c != 'c' {
			stack = append(stack, byte(c))
			continue
		}

		n := len(stack)
		if n < 2 || stack[n-2] != 'a' || stack[n-1] != 'b' {
			return false
		}
		stack = stack[:n-2]
	}

	return len(stack) == 0
}
