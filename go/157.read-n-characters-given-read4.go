package main

/*
 * @lc app=leetcode id=157 lang=golang
 *
 * [157] Read N Characters Given Read4
 *
 * https://leetcode.com/problems/read-n-characters-given-read4/description/
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		i := 0
		for n > 0 {
			if n-i >= 4 {
				x := read4(buf[i:])
				i += x
				if x < 4 {
					return i
				}
				continue
			}

			tmp := make([]byte, 4)
			x := read4(tmp)
			for j := 0; i < n && j < x; j++ {
				buf[i] = tmp[j]
				i++
			}
			break
		}

		return i
	}
}
