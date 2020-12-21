package main

// medium
// https://leetcode.com/problems/make-sum-divisible-by-p/

func minSubarray(nums []int, p int) int {
	d := 0
	for i, v := range nums {
		d = (d + v) % p
		nums[i] = d
	}

	if d == 0 {
		return 0
	}

	m := make(map[int]int, len(nums)+1)
	m[0] = -1
	best := len(nums)
	for i, v := range nums {
		if v >= d {
			j, ok := m[v-d]
			if ok && best > i-j {
				best = i - j
			}
		} else {
			j, ok := m[v+p-d]
			if ok && best > i-j {
				best = i - j
			}
		}

		m[v] = i
	}

	if best == len(nums) {
		return -1
	}
	return best
}
