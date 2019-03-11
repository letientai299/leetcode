package twosum

import "sort"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, x := range nums {
		if yIndex, ok := m[target-x]; ok {
			res := []int{i, yIndex}
			sort.Ints(res)
			return res
		}

		if _, ok := m[x]; !ok {
			m[x] = i
		}
	}
	return []int{}
}
