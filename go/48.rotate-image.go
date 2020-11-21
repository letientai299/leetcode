package main

func rotate(mat [][]int) {
	n := len(mat)
	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			mat[r][c], mat[r][n-1-c] = mat[r][n-1-c], mat[r][c]
		}
	}

	for r := 0; r < n-1; r++ {
		for c := 0; c < n-1-r; c++ {
			mat[r][c], mat[n-1-c][n-1-r] = mat[n-1-c][n-1-r], mat[r][c]
		}
	}
}
