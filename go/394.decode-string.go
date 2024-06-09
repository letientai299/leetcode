package main

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 *
 * https://leetcode.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (47.22%)
 * Total Accepted:    284.2K
 * Total Submissions: 544.7K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * Given an encoded string, return its decoded string.
 *
 * The encoding rule is: k[encoded_string], where the encoded_string inside the
 * square brackets is being repeated exactly k times. Note that k is guaranteed
 * to be a positive integer.
 *
 * You may assume that the input string is always valid; No extra white spaces,
 * square brackets are well-formed, etc.
 *
 * Furthermore, you may assume that the original data does not contain any
 * digits and that digits are only for those repeat numbers, k. For example,
 * there won't be input like 3a or 2[4].
 *
 *
 * Example 1:
 * Input: s = "3[a]2[bc]"
 * Output: "aaabcbc"
 * Example 2:
 * Input: s = "3[a2[c]]"
 * Output: "accaccacc"
 * Example 3:
 * Input: s = "2[abc]3[cd]ef"
 * Output: "abcabccdcdcdef"
 * Example 4:
 * Input: s = "abc3[cd]xyz"
 * Output: "abccdcdcdxyz"
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 30
 * s consists of lowercase English letters, digits, and square brackets
 * '[]'.
 * s is guaranteed to be a valid input.
 * All the integers in s are in the range [1, 300].
 *
 *
 */
func decodeString(s string) string {
	var dec func(s string) ([]byte, int)
	dec = func(s string) ([]byte, int) {
		mul := 0
		var res []byte

		for i := 0; i < len(s); i++ {
			c := s[i]
			switch {
			case '0' <= c && c <= '9':
				mul = 10*mul + int(c-'0')

			case c == '[':
				sub, steps := dec(s[i+1:])
				res = append(res, bytes.Repeat(sub, mul)...)
				mul = 0
				i += steps + 1

			case c == ']':
				return res, i

			default: // letter
				res = append(res, c)
				continue
			}
		}

		return res, len(s)
	}

	r, _ := dec(s)
	return string(r)
}

func decodeString_v1(s string) string {
	var dec func(i int) (string, int)

	dec = func(i int) (string, int) {
		if i >= len(s) {
			return "", i
		}

		var sb strings.Builder

		v := 0
		j := i
		for ; j < len(s); j++ {
			if '0' <= s[j] && s[j] <= '9' {
				v = 10*v + int(s[j]-'0')
				continue
			}

			if 'a' <= s[j] && s[j] <= 'z' {
				sb.WriteByte(s[j])
				continue
			}

			if s[j] == '[' {
				t, x := dec(j + 1)
				for k := 0; k < v; k++ {
					sb.WriteString(t)
				}
				j = x
				v = 0
				continue
			}

			if s[j] == ']' {
				j++
				break
			}
		}

		return sb.String(), j - 1
	}

	t, _ := dec(0)
	return t
}
