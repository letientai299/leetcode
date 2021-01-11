package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * medium
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 */
func combinationSum2(arr []int, n int) [][]int {
	sort.Ints(arr)
	var gen func(i, n int)
	res := make([][]int, 0)
	buf := make([]int, 0)

	gen = func(i, n int) {
		if i >= len(arr) || n < 0 {
			return
		}

		if n == 0 {
			b := make([]int, len(buf))
			copy(b, buf)
			res = append(res, b)
			return
		}

		l := len(buf)
		p := arr[i] - 1
		for i++; i < len(arr); i++ {
			if arr[i] != p && arr[i] <= n {
				buf = append(buf, arr[i])
				gen(i, n-arr[i])
				buf = buf[:l]
				p = arr[i]
			}
		}

		return
	}

	p := arr[0] - 1
	for i, x := range arr {
		if x == p {
			continue
		}
		p = x
		buf = append(buf, x)
		gen(i, n-x)
		buf = buf[:0]
	}

	return res
}
