package main

// medium
// https://leetcode.com/problems/spiral-matrix-iii/

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	up, down, left, right := r0, r0, c0, c0
	r, c := r0, c0
	x, y := 1, 0
	res := make([][]int, 0, R*C)
	turns := 0
	for len(res) != R*C {
		if turns == 0 {
			left--
			right++
			up--
			down++
		}
		for left <= c && c <= right && up <= r && r <= down {
			if 0 <= r && r < R && 0 <= c && c < C {
				res = append(res, []int{r, c})
			}
			c += x
			r += y
		}
		r -= y
		c -= x
		x, y = -y, x
		c += x
		r += y
		turns++
		turns %= 4
	}
	return res
}
