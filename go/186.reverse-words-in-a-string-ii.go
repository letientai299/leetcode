package main

// Reverse Words in a String II
//
// Medium
//
// https://leetcode.com/problems/reverse-words-in-a-string-ii/
func reverseWords(s []byte) {
	n := len(s)
	rev := func(start, end int) {
		for i := 0; i < (end-start)/2; i++ {
			s[start+i], s[end-1-i] = s[end-1-i], s[start+i]
		}
	}

	rev(0, n)
	start := 0
	for end := 1; end < n; end++ {
		if s[end] == ' ' {
			rev(start, end)
			start = end + 1
		}
	}
	rev(start, n)
}
