package main

import (
	"sort"
)

// Find Kth Largest XOR Coordinate Value
//
// Medium
//
// https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/
//
// You are given a 2D `matrix` of size `m x n`, consisting of non-negative
// integers. You are also given an integer `k`.
//
// The **value** of coordinate `(a, b)` of the matrix is the XOR of all
// `matrix[i][j]` where `0 <= i <= a < m` and `0 <= j <= b < n` **(0-indexed)**.
//
// Find the `kth` largest value **(1-indexed)** of all the coordinates of
// `matrix`.
//
// **Example 1:**
//
// ```
// Input: matrix = [[5,2],[1,6]], k = 1
// Output: 7
// Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the
// largest value.
// ```
//
// **Example 2:**
//
// ```
// Input: matrix = [[5,2],[1,6]], k = 2
// Output: 5
// Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest
// value.
// ```
//
// **Example 3:**
//
// ```
// Input: matrix = [[5,2],[1,6]], k = 3
// Output: 4
// Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd
// largest value.
// ```
//
// **Example 4:**
//
// ```
// Input: matrix = [[5,2],[1,6]], k = 4
// Output: 0
// Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which
// is the 4th largest value.
// ```
//
// **Constraints:**
//
// - `m == matrix.length`
// - `n == matrix[i].length`
// - `1 <= m, n <= 1000`
// - `0 <= matrix[i][j] <= 106`
// - `1 <= k <= m * n`
func kthLargestValue(mat [][]int, k int) int {
	m := len(mat)
	n := len(mat[0])
	all := make([]int, m*n)
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			v := mat[y][x]
			if x > 0 {
				v ^= mat[y][x-1]
			}
			if y > 0 {
				v ^= mat[y-1][x]
			}

			if y > 0 && x > 0 {
				v ^= mat[y-1][x-1]
			}

			mat[y][x] = v
			all[y*n+x] = v
		}
	}

	sort.Ints(all)
	return all[len(all)-k]
}
