package main

// medium

// TODO (tai): can be faster, 30%
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	for i := range nums[1:] {
		if nums[i] == 0 && nums[i+1] == 0 {
			return true
		}
	}

	if k == 0 {
		return false
	}

	if k < 0 {
		k = -k
	}

	m := make(map[int]int)

	s := 0
	for i := 0; i < n; i++ {
		s += nums[i]
		m[s] = i
	}

	if k == s || s == 0 || k == 1 {
		return true
	}

	if k > s {
		return false
	}

	m[0] = -1

	mul := s / k
	for v, i := range m {
		for x := 1; x <= mul; x++ {
			if k*x > v {
				j, ok := m[k*x+v]
				if ok && j-i >= 2 {
					return true
				}
			} else {
				j, ok := m[v-k*x]
				if ok && i-j >= 2 {
					return true
				}
			}
		}
	}
	return false
}
