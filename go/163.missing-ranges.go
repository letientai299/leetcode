package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=163 lang=golang
 *
 * [163] Missing Ranges
 *
 * https://leetcode.com/problems/missing-ranges/description/
 */
func findMissingRanges(nums []int, lower int, upper int) []string {
	fmtRange := func(a ...int) string {
		if len(a) == 1 || a[0] == a[1] {
			return strconv.Itoa(a[0])
		}
		return strconv.Itoa(a[0]) + "->" + strconv.Itoa(a[1])
	}

	if len(nums) == 0 {
		return []string{fmtRange(lower, upper)}
	}

	var res []string

	a := lower - 1
	for i := 0; i < len(nums); i++ {
		b := nums[i]
		if a < b-1 {
			res = append(res, fmtRange(a+1, b-1))
		}
		a = b
	}

	if a < upper {
		res = append(res, fmtRange(a+1, upper))
	}
	return res
}
