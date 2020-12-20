package main

import "sort"

// medium

func maximumUniqueSubarray(nums []int) int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	m := make(map[int][]int)

	pre := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		sum[i] = v + pre
		pre = sum[i]
		m[v] = append(m[v], i)
	}

	best := nums[0]
	start := 0
	end := 1

	for ; end < len(nums); end++ {
		v := nums[end]
		locs := m[v]
		x := sort.SearchInts(locs, start)
		y := sort.SearchInts(locs, end)
		if x == y {
			// not duplicated yet
			continue
		}

		r := sum[end] - sum[start] - nums[end] + nums[start]
		start = locs[x] + 1
		if r > best {
			best = r
		}
	}

	v := sum[end-1] - sum[start] + nums[start]
	if v > best {
		best = v
	}

	return best
}
