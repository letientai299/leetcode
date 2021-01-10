package main

// medium
// 1664. Ways to Make a Fair Array

func waysToMakeFair(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	var right [2]int
	for i, v := range nums {
		right[i%2] += v
	}

	var left [2]int
	r := 0
	for i, v := range nums {
		j := i % 2
		right[j] -= v
		if right[j]+left[1-j] == right[1-j]+left[j] {
			r++
		}
		left[j] += v
	}

	return r
}
