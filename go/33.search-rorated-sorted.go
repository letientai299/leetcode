package main

import (
	"sort"
)

// medium
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, t int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		switch {
		case t == nums[l]:
			return l
		case t == nums[r]:
			return r
		case t == nums[m]:
			return m

		case nums[l] <= nums[r]:
			i := sort.SearchInts(nums[l:r], t)
			if i == r-l || nums[i+l] != t {
				return -1
			}
			return i + l

		case t < nums[r] && nums[m+1] <= t:
			i := sort.SearchInts(nums[m+1:r], t)
			if i == r-m-1 || nums[i+m+1] != t {
				return -1
			}
			return i + m + 1

		case nums[l] < t && t < nums[m]:
			i := sort.SearchInts(nums[l:m], t)
			if i == m-l || nums[i+l] != t {
				return -1
			}
			return i + l
		case nums[l] > nums[m]:
			r = m - 1
		case nums[r] < nums[m]:
			l = m + 1
		default:
			panic("wtf")
		}
	}

	return -1
}
