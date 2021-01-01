package main

import "sort"

func frequencySort1636(nums []int) []int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}

	var tmp [][]int
	for k, v := range m {
		tmp = append(tmp, []int{k, v})
	}

	sort.Slice(tmp, func(i, j int) bool {
		a := tmp[i]
		b := tmp[j]
		if a[1] < b[1] {
			return true
		}

		if a[1] == b[1] {
			return a[0] > b[0]
		}

		return false
	})

	for i := 0; i < len(nums); {
		for _, a := range tmp {
			for ; a[1] > 0; a[1]-- {
				nums[i] = a[0]
				i++
			}
		}
	}
	return nums
}
