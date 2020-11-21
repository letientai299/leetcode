package main

// medium
// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	all := 1
	for i := 2; i <= len(nums); i++ {
		all *= i
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
