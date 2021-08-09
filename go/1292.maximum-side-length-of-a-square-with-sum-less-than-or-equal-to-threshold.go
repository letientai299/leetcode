package main

// Maximum Side Length of a Square with Sum Less than or Equal to Threshold
//
// Medium
//
// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
//
// Given a `m x n` matrix `mat` and an integer `threshold`. Return the maximum
// side-length of a square with a sum less than or equal to `threshold` or
// return **0** if there is no such square.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/12/05/e1.png)
//
// ```
// Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
// Output: 2
// Explanation: The maximum side length of square with sum less than 4 is 2 as
// shown.
//
// ```
//
// **Example 2:**
//
// ```
// Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]],
// threshold = 1
// Output: 0
//
// ```
//
// **Example 3:**
//
// ```
// Input: mat = [[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]], threshold = 6
// Output: 3
//
// ```
//
// **Example 4:**
//
// ```
// Input: mat = [[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]],
// threshold = 40184
// Output: 2
//
// ```
//
// **Constraints:**
//
// - `1 <= m, n <= 300`
// - `m == mat.length`
// - `n == mat[i].length`
// - `0 <= mat[i][j] <= 10000`
// - `0 <= threshold <= 10^5`
func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	best := 0

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			up := 0
			if y > 0 {
				up = mat[y-1][x]
			}
			left := 0
			if x > 0 {
				left = mat[y][x-1]
			}
			upLeft := 0
			if y > 0 && x > 0 {
				upLeft = mat[y-1][x-1]
			}

			if best == 0 && mat[y][x] <= threshold {
				best = 1
			}

			mat[y][x] += up + left - upLeft
			if best == 0 {
				continue
			}

			for i := best + 1; x-i >= -1 && y-i >= -1; i++ {
				v := mat[y][x]
				if x-i >= 0 {
					v -= mat[y][x-i]
				}

				if y-i >= 0 {
					v -= mat[y-i][x]
				}

				if y-i >= 0 && x-i >= 0 {
					v += mat[y-i][x-i]
				}

				if v > threshold {
					break
				}

				best = i
			}
		}
	}

	return best
}
