package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	soldies := make([]int, 0, len(mat))
	m := make(map[int][]int)
	for id, r := range mat {
		s := 0
		for _, x := range r {
			if x == 0 {
				break
			}
			s++
		}
		if _, ok := m[s]; !ok {
			soldies = append(soldies, s)
		}
		m[s] = append(m[s], id)
	}

	sort.Ints(soldies)
	var res []int
	for _, x := range soldies {
		res = append(res, m[x]...)
		if len(res) >= k {
			break
		}
	}

	return res[:k]
}
