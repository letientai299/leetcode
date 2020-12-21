package main

/*
 * @lc app=leetcode id=1253 lang=golang
 *
 * [1253] Reconstruct a 2-Row Binary Matrix
 *
 * https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/description/
 *
 * algorithms
 * Medium (37.77%)
 * Total Accepted:    12.9K
 * Total Submissions: 31.3K
 * Testcase Example:  '2\n1\n[1,1,1]'
 *
 * Given the following details of a matrix with n columns and 2 rows :
 *
 *
 * The matrix is a binary matrix, which means each element in the matrix can be
 * 0 or 1.
 * The sum of elements of the 0-th(upper) row is given as upper.
 * The sum of elements of the 1-st(lower) row is given as lower.
 * The sum of elements in the i-th column(0-indexed) is colsum[i], where colsum
 * is given as an integer array with length n.
 *
 *
 * Your task is to reconstruct the matrix with upper, lower and colsum.
 *
 * Return it as a 2-D integer array.
 *
 * If there are more than one valid solution, any of them will be accepted.
 *
 * If no valid solution exists, return an empty 2-D array.
 *
 *
 * Example 1:
 *
 *
 * Input: upper = 2, lower = 1, colsum = [1,1,1]
 * Output: [[1,1,0],[0,0,1]]
 * Explanation: [[1,0,1],[0,1,0]], and [[0,1,1],[1,0,0]] are also correct
 * answers.
 *
 *
 * Example 2:
 *
 *
 * Input: upper = 2, lower = 3, colsum = [2,2,1,1]
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: upper = 5, lower = 5, colsum = [2,1,2,0,1,0,1,2,0,1]
 * Output: [[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= colsum.length <= 10^5
 * 0 <= upper, lower <= colsum.length
 * 0 <= colsum[i] <= 2
 *
 *
 */
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	s := 0
	for _, v := range colsum {
		s += v
	}

	if s != upper+lower {
		return make([][]int, 0)
	}

	n := len(colsum)
	mat := make([][]int, 2)
	mat[0] = colsum
	mat[1] = make([]int, n)
	for i, v := range colsum {
		if upper < 0 || lower < 0 {
			return make([][]int, 0)
		}

		switch v {
		case 0:
			mat[0][i] = 0
			mat[1][i] = 0
		case 2:
			mat[0][i] = 1
			mat[1][i] = 1
			upper--
			lower--
		case 1:
			if upper >= lower {
				mat[0][i] = 1
				mat[1][i] = 0
				upper--
			} else {
				mat[1][i] = 1
				mat[0][i] = 0
				lower--
			}
		}
	}

	if upper != 0 || lower != 0 {
		return make([][]int, 0)
	}

	return mat
}
