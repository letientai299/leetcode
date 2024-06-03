package main

// Find Permutation
//
// # Medium
//
// https://leetcode.com/problems/find-permutation/
func findPermutation(s string) []int {
	n := len(s) + 1 // len of result array
	res := make([]int, 0, n)
	res = append(res, 1)

	cnt := 1
	k := 0
	for i := 1; i < len(s)+1; i++ {
		if i < len(s) && s[i] == s[i-1] {
			cnt++
			continue
		}

		if s[i-1] == 'I' {
			for j := 1; j <= cnt; j++ {
				res = append(res, res[k]+j)
			}
			cnt = 1
			continue
		}

		k = len(res) - 1
		res[k] += cnt
		for j := 1; j <= cnt; j++ {
			res = append(res, res[k]-j)
		}
		cnt = 1
	}

	return res
}
