package main

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := countBitsSWAR(arr[i])
		b := countBitsSWAR(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

// https://stackoverflow.com/a/109025
func countBitsSWAR(i int) int {
	n := uint32(i)
	n = n - ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = ((n + (n >> 4)) & 0x0F0F0F0F) * 0x01010101
	return int(n >> 24)
}
