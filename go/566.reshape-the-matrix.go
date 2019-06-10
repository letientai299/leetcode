package main

/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 *
 * https://leetcode.com/problems/reshape-the-matrix/description/
 *
 * algorithms
 * Easy (58.91%)
 * Total Accepted:    78K
 * Total Submissions: 132.4K
 * Testcase Example:  '[[1,2],[3,4]]\n1\n4'
 *
 * In MATLAB, there is a very useful function called 'reshape', which can
 * reshape a matrix into a new one with different size but keep its original
 * data.
 *
 *
 *
 * You're given a matrix represented by a two-dimensional array, and two
 * positive integers r and c representing the row number and column number of
 * the wanted reshaped matrix, respectively.
 *
 * ⁠The reshaped matrix need to be filled with all the elements of the original
 * matrix in the same row-traversing order as they were.
 *
 *
 *
 * If the 'reshape' operation with given parameters is possible and legal,
 * output the new reshaped matrix; Otherwise, output the original matrix.
 *
 *
 * Example 1:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 1, c = 4
 * Output:
 * [[1,2,3,4]]
 * Explanation:The row-traversing of nums is [1,2,3,4]. The new reshaped matrix
 * is a 1 * 4 matrix, fill it row by row by using the previous list.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 2, c = 4
 * Output:
 * [[1,2],
 * ⁠[3,4]]
 * Explanation:There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So
 * output the original matrix.
 *
 *
 *
 * Note:
 *
 * The height and width of the given matrix is in range [1, 100].
 * The given r and c are all positive.
 *
 *
 */
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums[0]) == 0 {
		return nums
	}

	m := len(nums)
	n := len(nums[0])
	total := m * n
	if total != r*c {
		return nums
	}

	var res [][]int

	i, j := 0, 0
	for y := 0; y < r; y++ {
		var row []int
		for x := 0; x < c; x++ {
			row = append(row, nums[i][j])
			j++
			if j >= n {
				j = 0
				i++
			}
		}
		res = append(res, row)
	}

	/*
	// This solution use less memory, but take more time, perhaps each
	// allocation take more time than before
	var res [][]int
	var remain []int

	for _, row := range nums {
		remain = append(remain, row...)
		if c > len(remain) {
			continue // keep appending
		}
		for len(remain) > c {
			next := remain[:c]
			remain = remain[c:]
			res = append(res, next)
		}
	}

	if len(remain) != 0 {
		res = append(res, remain)
	}
	*/
	return res
}
