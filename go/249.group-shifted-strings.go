package main

// Group Shifted Strings
//
// Medium
//
// https://leetcode.com/problems/group-shifted-strings/
func groupStrings(strings []string) [][]string {
	index := func(s string) int {
		anchor := int32(s[0])
		v := len(s)
		for _, c := range s[1:] {
			d := c - anchor
			if d < 0 {
				d += 26
			}
			v = 26*v + int(d)
		}

		return v
	}

	m := make(map[int][]string)
	for _, s := range strings {
		i := index(s)
		m[i] = append(m[i], s)
	}

	r := make([][]string, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
