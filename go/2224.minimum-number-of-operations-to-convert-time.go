package main

// Minimum Number of Operations to Convert Time
//
// Easy
//
// https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/
//
// You are given two strings `current` and `correct` representing two **24-hour
// times**.
//
// 24-hour times are formatted as `"HH:MM"`, where `HH` is between `00` and
// `23`, and `MM` is between `00` and `59`. The earliest 24-hour time is
// `00:00`, and the latest is `23:59`.
//
// In one operation you can increase the time `current` by `1`, `5`, `15`, or
// `60` minutes. You can perform this operation **any** number of times.
//
// Return _the **minimum number of operations** needed to convert_ `current`
// _to_ `correct`.
//
// **Example 1:**
//
// ```
// Input: current = "02:30", correct = "04:35"
// Output: 3
// Explanation:
// We can convert current to correct in 3 operations as follows:
// - Add 60 minutes to current. current becomes "03:30".
// - Add 60 minutes to current. current becomes "04:30".
// - Add 5 minutes to current. current becomes "04:35".
// It can be proven that it is not possible to convert current to correct in
// fewer than 3 operations.
// ```
//
// **Example 2:**
//
// ```
// Input: current = "11:00", correct = "11:01"
// Output: 1
// Explanation: We only have to add one minute to current, so the minimum number
// of operations needed is 1.
//
// ```
//
// **Constraints:**
//
// - `current` and `correct` are in the format `"HH:MM"`
// - `current <= correct`
func convertTime(current string, correct string) int {
	ch := int((current[0]-'0')*10 + current[1] - '0') // current hour
	cm := int((current[3]-'0')*10 + current[4] - '0')

	wh := int((correct[0]-'0')*10 + correct[1] - '0')
	wm := int((correct[3]-'0')*10 + correct[4] - '0')

	r := 0
	if wh >= ch {
		r += wh - ch
	} else {
		r += 12 - ch + wh
	}

	if wm < cm {
		wm += 60 - cm
		// we're turning the clock to 1 more hour, thus, remove the hour diff original
		r--
	} else {
		wm -= cm
	}

	r += wm / 15
	wm %= 15
	r += wm / 5
	wm %= 5
	r += wm
	return r
}
