package main

/*
* @lc app=leetcode id=434 lang=golang
*
* [434] Number of Segments in a String
*
* https://leetcode.com/problems/number-of-segments-in-a-string/description/
*
* algorithms
* Easy (36.77%)
* Total Accepted:    52.9K
* Total Submissions: 144K Testcase Example:  '"Hello, my name is John"'
*
* Count the number of segments in a string, where a segment is defined to be a
* contiguous sequence of non-space characters.
*
* Please note that the string does not contain any non-printable characters.
*
* Example:
*
* Input: "Hello, my name is John"
* Output: 5
*
*
 */
func countSegments(s string) int {
	n := 0
	prev := ' '
	for _, c := range s {
		if c == rune(' ') {
			if prev != rune(' ') {
				n++
			}
		}
		prev = c
	}
	if prev != rune(' ') {
		n++
	}
	return n
}
