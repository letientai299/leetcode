package main

// medium
// https://leetcode.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}
	factorial := func(i int) int {
		if i < 2 {
			return 1
		}
		n := 1
		for x := 2; x <= i; x++ {
			n *= x
		}
		return n
	}

	all := factorial(len(nums))
	for _, v := range m {
		all /= factorial(v)
	}

	res := make([][]int, 0, all)
	all--
	t := make([]int, len(nums))
	copy(t, nums)
	res = append(res, t)
	for all > 0 {
		nextPermutation(nums)
		t := make([]int, len(nums))
		copy(t, nums)
		res = append(res, t)
		all--
	}
	return res
}
