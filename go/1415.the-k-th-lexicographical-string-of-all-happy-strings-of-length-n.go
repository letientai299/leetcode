package main

// The k-th Lexicographical String of All Happy Strings of Length n
//
// Medium
//
// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
//
// A **happy string** is a string that:
//
// - consists only of letters of the set `['a', 'b', 'c']`.
// - `s[i] != s[i + 1]` for all values of `i` from `1` to `s.length - 1` (string
// is 1-indexed).
//
// For example, strings **"abc", "ac", "b"** and **"abcbabcbcb"** are all happy
// strings and strings **"aa", "baa"** and **"ababbc"** are not happy strings.
//
// Given two integers `n` and `k`, consider a list of all happy strings of
// length `n` sorted in lexicographical order.
//
// Return _the kth string_ of this list or return an **empty string** if there
// are less than `k` happy strings of length `n`.
//
// **Example 1:**
//
// ```
// Input: n = 1, k = 3
// Output: "c"
// Explanation: The list ["a", "b", "c"] contains all happy strings of length 1.
// The third string is "c".
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1, k = 4
// Output: ""
// Explanation: There are only 3 happy strings of length 1.
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 3, k = 9
// Output: "cab"
// Explanation: There are 12 different happy string of length 3 ["aba", "abc",
// "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]. You
// will find the 9th string = "cab"
//
// ```
//
// **Example 4:**
//
// ```
// Input: n = 2, k = 7
// Output: ""
//
// ```
//
// **Example 5:**
//
// ```
// Input: n = 10, k = 100
// Output: "abacbabacb"
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 10`
// - `1 <= k <= 100`
func getHappyString(n int, k int) string {
	// 3 choices for the first letter,
	// then 2 choices for n-1 remained letter
	total := 3 << (n - 1)

	if k > total {
		return ""
	}

	k--
	total /= 3

	bs := make([]byte, n)
	bs[0] = byte(k/total + 'a')

	k %= total
	for i := 1; i < n; i++ {
		total /= 2
		index := k / total
		k%=total

		c := 'a'
		switch bs[i-1] {
		case 'a':
			if index == 0 {
				c = 'b'
			} else {
				c = 'c'
			}
		case 'b':
			if index == 0 {
				c = 'a'
			} else {
				c = 'c'
			}
		case 'c':
			if index == 0 {
				c = 'a'
			} else {
				c = 'b'
			}
		}
		bs[i] = byte(c)
	}
	return string(bs)
}
