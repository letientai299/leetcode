package main

/*
* @lc app = leetcode id = 118 lang =golang
*
* [118] Pascal's Triangle
*
* https://leetcode.com/problems/pascals-triangle/description/
*
* algorithms
* Easy (44.81%)
* Total Accepted:    232.9K
* Total Submissions: 519.6K
* Testcase Example:  '5'
*
* Given a non-negative integer numRows, generate the first numRows of Pascal's
* triangle.
*
*
* In Pascal's triangle, each number is the sum of the two numbers directly
* above it.
*
* Example:
*
*
* Input: 5
* Output:
* [
* ⁠    [1],
* ⁠   [1, 1],
* ⁠  [1, 2, 1],
* ⁠ [1, 3, 3,1],
* ⁠[1, 4,6, 4, 1]
* ]
*
*
 */
func generate(n int) [][]int {
	if n == 0 {
		return nil
	}

	res := [][]int{{1}}
	for i := 1; i < n; i++ {
		next := []int{1}
		for x := 0; x < i-1; x++ {
			next = append(next, res[i-1][x+1]+res[i-1][x])
		}
		next = append(next, 1)
		res = append(res, next)
	}
	return res
}
