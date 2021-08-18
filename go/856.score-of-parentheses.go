package main

// Score of Parentheses
//
// Medium
//
// https://leetcode.com/problems/score-of-parentheses/
//
// Given a balanced parentheses string `s`, return _the **score** of the
// string_.
//
// The **score** of a balanced parentheses string is based on the following
// rule:
//
// - `"()"` has score `1`.
// - `AB` has score `A + B`, where `A` and `B` are balanced parentheses strings.
// - `(A)` has score `2 * A`, where `A` is a balanced parentheses string.
//
// **Example 1:**
//
// ```
// Input: s = "()"
// Output: 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "(())"
// Output: 2
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "()()"
// Output: 2
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = "(()(()))"
// Output: 6
//
// ```
//
// **Constraints:**
//
// - `2 <= s.length <= 50`
// - `s` consists of only `'('` and `')'`.
// - `s` is a balanced parentheses string.
func scoreOfParentheses(s string) int {
	var stacks []int
	var sums []int

	pre := ')'
	for _, c := range s {
		n := len(stacks)
		if c == '(' {
			if pre != c {
				stacks = append(stacks, 1)
				sums = append(sums, 0)
			} else {
				stacks[n-1]++
			}

			pre = c
			continue
		}

		pre = c
		stacks[n-1]--
		if sums[n-1] == 0 {
			sums[n-1] = 1
		} else {
			sums[n-1] *= 2
		}

		if stacks[n-1] == 0 && n > 1 {
			sums[n-2] += sums[n-1]
			stacks = stacks[:n-1]
			sums = sums[:n-1]
		}
	}

	return sums[0]
}
