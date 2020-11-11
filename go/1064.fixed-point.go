package main

// https://leetcode.com/problems/fixed-point/

func fixedPoint(a []int) int {
	for i, x := range a {
		if i == x {
			return i
		}
	}
	return -1
}
