package main

func diagonalSum(mat [][]int) int {
	n := len(mat)
	s := 0
	for i := 0; i < n; i++ {
		s = s + mat[i][i] + mat[i][n-1-i]
	}

	if n%2 == 1 {
		s -= mat[n/2][n/2]
	}
	return s
}
