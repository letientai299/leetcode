package main

// medium

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	b, e := 0, 1
	sum := nums[b]
	if sum >= s {
		return 1
	}

	min := n + 1

	for e < n {
		sum += nums[e]
		for sum < s && e < n {
			e++
			if e == n {
				e--
				break
			}
			sum += nums[e]
		}
		for b <= e && sum-nums[b] >= s {
			sum -= nums[b]
			b++
		}

		if sum >= s && min > e-b+1 {
			min = e - b + 1
		}
		e++
	}

	if sum >= s && min > e-b+1 {
		min = e - b + 1
	}

	if min > n {
		return 0
	}

	return min
}
