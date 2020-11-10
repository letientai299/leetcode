package main

func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	res := 0
	for _, v := range m {
		res += v * (v - 1) / 2
	}
	return res
}
