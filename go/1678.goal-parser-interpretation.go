package main

import "strings"

// Goal Parser Interpretation
//
// Easy
//
// https://leetcode.com/problems/goal-parser-interpretation/
//
// You own a **Goal Parser** that can interpret a string `command`. The
// `command` consists of an alphabet of `"G"`, `"()"` and/or `"(al)"` in some
// order. The Goal Parser will interpret `"G"` as the string `"G"`, `"()"` as
// the string `"o"`, and `"(al)"` as the string `"al"`. The interpreted strings
// are then concatenated in the original order.
//
// Given the string `command`, return _the **Goal Parser**'s interpretation of_
// `command`.
//
// **Example 1:**
//
// ```
// Input: command = "G()(al)"
// Output: "Goal"
// Explanation: The Goal Parser interprets the command as follows:
// G -> G
// () -> o
// (al) -> al
// The final concatenated result is "Goal".
//
// ```
//
// **Example 2:**
//
// ```
// Input: command = "G()()()()(al)"
// Output: "Gooooal"
//
// ```
//
// **Example 3:**
//
// ```
// Input: command = "(al)G(al)()()G"
// Output: "alGalooG"
//
// ```
//
// **Constraints:**
//
// - `1 <= command.length <= 100`
// - `command` consists of `"G"`, `"()"`, and/or `"(al)"` in some order.
func interpret(cmd string) string {
	var sb strings.Builder
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case 'G':
			sb.WriteByte('G')
		case '(':
			if cmd[i+1] == 'a' {
				sb.WriteString("al")
				i += 3
			} else {
				sb.WriteByte('o')
				i++
			}
		}
	}
	return sb.String()
}
