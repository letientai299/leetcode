package main

// medium

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	s := 0
	for i := 1; i < n; i++ {
		s += nums[i] - nums[0]
	}

	v := nums[0]
	nums[0] = s
	d := 0

	for i := 1; i < n; i++ {
		v, d = nums[i], nums[i]-v
		nums[i] = nums[i-1] - (n-2*i)*d
	}

	return nums
}
