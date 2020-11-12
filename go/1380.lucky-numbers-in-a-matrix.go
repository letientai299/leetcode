package main

func luckyNumbers(mat [][]int) []int {
	var rowMins [][]int
	for _, r := range mat {
		m := 1 << 31
		id := 0
		for i, c := range r {
			if m > c {
				m = c
				id = i
			}
		}
		rowMins = append(rowMins, []int{m, id})
	}

	var res []int
	for _, r := range rowMins {
		v := r[0]
		c := r[1]
		yes:= true
		for i := 0; i < len(mat); i++ {
			if mat[i][c] > v {
				yes= false
				break
			}
		}
		if yes{
			res = append(res, v)
		}
	}
	return res
}
