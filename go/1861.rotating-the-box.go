package main

// Rotating the Box
//
// Medium
//
// https://leetcode.com/problems/rotating-the-box/
//
// You are given an `m x n` matrix of characters `box` representing a side-view
// of a box. Each cell of the box is one of the following:
//
// - A stone `'#'`
// - A stationary obstacle `'*'`
// - Empty `'.'`
//
// The box is rotated **90 degrees clockwise**, causing some of the stones to
// fall due to gravity. Each stone falls down until it lands on an obstacle,
// another stone, or the bottom of the box. Gravity **does not** affect the
// obstacles' positions, and the inertia from the box's rotation **does not**
// affect the stones' horizontal positions.
//
// It is **guaranteed** that each stone in `box` rests on an obstacle, another
// stone, or the bottom of the box.
//
// Return _an_ `n x m` _matrix representing the box after the rotation described
// above_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcodewithstones.png)
//
// ```
// Input: box = [["#",".","#"]]
// Output: [["."],
//          ["#"],
//          ["#"]]
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode2withstones.png)
//
// ```
// Input: box = [["#",".","*","."],
//               ["#","#","*","."]]
// Output: [["#","."],
//          ["#","#"],
//          ["*","*"],
//          [".","."]]
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode3withstone.png)
//
// ```
// Input: box = [["#","#","*",".","*","."],
//               ["#","#","#","*",".","."],
//               ["#","#","#",".","#","."]]
// Output: [[".","#","#"],
//          [".","#","#"],
//          ["#","#","*"],
//          ["#","*","."],
//          ["#",".","*"],
//          ["#",".","."]]
//
// ```
//
// **Constraints:**
//
// - `m == box.length`
// - `n == box[i].length`
// - `1 <= m, n <= 500`
// - `box[i][j]` is either `'#'`, `'*'`, or `'.'`.
func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
	}

	for x, row := range box {
		i := n - 1
		for i >= 0 && (row[i] == '*' || row[i] == '#') {
			i--
		}

		for j := i - 1; i >= 0 && j >= 0; j-- {
			if row[j] == '*' {
				for i = j - 1; i >= 0 && (row[i] == '*' || row[i] == '#'); i-- {
				}
				j = i
				continue
			}

			if row[j] == '#' {
				row[i], row[j] = row[j], row[i]
				i--
			}
		}

		for y, b := range row {
			res[y][m-1-x] = b
		}
	}

	return res
}
