package main

// https: //leetcode.com/problems/find-anagram-mappings/
func anagramMappings(a []int, b []int) []int {
	m := make(map[int][]int, len(b))
	for i, x := range b {
		if _, ok := m[x]; ok {
			m[x] = append(m[x], i)
		} else {
			m[x] = []int{i}
		}
	}

	p := make([]int, len(a))
	for i, x := range a {
		ids := m[x]
		p[i] = ids[0]
		m[x] = ids[1:]
	}
	return p
}
