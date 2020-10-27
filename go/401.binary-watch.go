package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 *
 * https://leetcode.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (46.01%)
 * Total Accepted:    88.2K
 * Total Submissions: 184.1K
 * Testcase Example:  '0'
 *
 * A binary watch has 4 LEDs on the top which represent the hours (0-11), and
 * the 6 LEDs on the bottom represent the minutes (0-59).
 * Each LED represents a zero or one, with the least significant bit on the
 * right.
 *
 * For example, the above binary watch reads "3:25".
 *
 * Given a non-negative integer n which represents the number of LEDs that are
 * currently on, return all possible times the watch could represent.
 *
 * Example:
 * Input: n = 1
 * Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04",
 * "0:08", "0:16", "0:32"]
 *
 *
 * Note:
 *
 * The order of output does not matter.
 * The hour must not contain a leading zero, for example "01:00" is not valid,
 * it should be "1:00".
 * The minute must be consist of two digits and may contain a leading zero, for
 * example "10:2" is not valid, it should be "10:02".
 *
 *
 */
func readBinaryWatch(n int) []string {
	var res []string
	nums := genNum(n, 10)

	for _, x := range nums {
		h := (x & (15 << 6)) >> 6
		if h >= 12 {
			continue // hour is only 0 to 11
		}
		m := x & (1<<6 - 1)

		if m >= 60 {
			continue
		}
		res = append(res, fmt.Sprintf("%d:%02d", h, m))
	}
	return res
}

func genNum(n int, slot uint) []int {
	var res []int
	v := 0

	var recur func(i int, pos uint)

	recur = func(bi int, pos uint) {
		if bi >= n {
			res = append(res, v)
			return
		}

		for x := pos; x < slot-uint(n-bi)+1; x++ {
			v += 1 << x
			recur(bi+1, x+1)
			v -= 1 << x
		}
	}

	recur(0, 0)
	return res
}
