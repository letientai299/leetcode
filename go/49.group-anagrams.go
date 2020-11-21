package main

// medium
// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(ss []string) [][]string {
	m := make(map[int][]string)
	base := [26]int{}
	base[0] = 1
	for c := 'b'; c <= 'z'; c++ {
		base[c-'a'] = base[c-'a'-1] * 26
	}
	for _, s := range ss {
		h := 0
		for _, c := range s {
			h += base[int(c-'a')]
		}
		m[h] = append(m[h], s)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
