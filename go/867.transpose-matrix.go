package main

// https://leetcode.com/problems/transpose-matrix/
func transpose(arr [][]int) [][]int {
	var res [][]int
	for _, a := range arr {
		for y, c := range a {
			if y >= len(res) {
				res = append(res, []int{})
			}
			res[y] = append(res[y], c)
		}
	}
	return res
}
