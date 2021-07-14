package main

import "fmt"

// Latest Time by Replacing Hidden Digits
//
// Easy
//
// https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/
//
// You are given a string `time` in the form of ` hh:mm`, where some of the
// digits in the string are hidden (represented by `?`).
//
// The valid times are those inclusively between `00:00` and `23:59`.
//
// Return _the latest valid time you can get from_ `time` _by replacing the
// hidden_ _digits_.
//
// **Example 1:**
//
// ```
// Input: time = "2?:?0"
// Output: "23:50"
// Explanation: The latest hour beginning with the digit '2' is 23 and the
// latest minute ending with the digit '0' is 50.
//
// ```
//
// **Example 2:**
//
// ```
// Input: time = "0?:3?"
// Output: "09:39"
//
// ```
//
// **Example 3:**
//
// ```
// Input: time = "1?:22"
// Output: "19:22"
//
// ```
//
// **Constraints:**
//
// - `time` is in the format `hh:mm`.
// - It is guaranteed that you can produce a valid time from the given string.
func maximumTime(time string) string {
	a, b, c, d := time[0], time[1], time[3], time[4]
	h := 0
	if a == '?' && b == '?' {
		h = 23
	} else if a == '?' {
		h = int(b - '0')
		if h > 3 {
			h += 10
		} else {
			h += 20
		}
	} else if b == '?' {
		h = 10 * int(a-'0')
		if h < 20 {
			h += 9
		} else {
			h += 3
		}
	} else {
		h = int((a-'0')*10 + b - '0')
	}

	m := 0
	if c == '?' && d == '?' {
		m = 59
	} else if c == '?' {
		m = 50 + int(d-'0')
	} else if d == '?' {
		m = 9 + 10*int(c-'0')
	} else {
		m = int((c-'0')*10 + d - '0')
	}

	return fmt.Sprintf("%02d:%02d", h, m)
}
