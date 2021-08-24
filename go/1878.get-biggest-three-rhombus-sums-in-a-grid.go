package main

import "sort"

// Get Biggest Three Rhombus Sums in a Grid
//
// Medium
//
// https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/
//
// You are given an `m x n` integer matrix `grid`​​​.
//
// A **rhombus sum** is the sum of the elements that form **the** **border** of
// a regular rhombus shape in `grid`​​​. The rhombus must have the shape of a
// square rotated 45 degrees with each of the corners centered in a grid cell.
// Below is an image of four valid rhombus shapes with the corresponding colored
// cells that should be included in each **rhombus sum**:
//
// ![](https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-desc-2.png)
//
// Note that the rhombus can have an area of 0, which is depicted by the purple
// rhombus in the bottom right corner.
//
// Return _the biggest three **distinct rhombus sums** in the_ `grid` _in
// **descending order**_ _. If there are less than three distinct values, return
// all of them_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex1.png)
//
// ```
// Input: grid =
// [[3,4,5,1,3],[3,3,4,2,3],[20,30,200,40,10],[1,5,5,4,1],[4,3,2,2,5]]
// Output: [228,216,211]
// Explanation: The rhombus shapes for the three biggest distinct rhombus sums
// are depicted above.
// - Blue: 20 + 3 + 200 + 5 = 228
// - Red: 200 + 2 + 10 + 4 = 216
// - Green: 5 + 200 + 4 + 2 = 211
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex2.png)
//
// ```
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [20,9,8]
// Explanation: The rhombus shapes for the three biggest distinct rhombus sums
// are depicted above.
// - Blue: 4 + 2 + 6 + 8 = 20
// - Red: 9 (area 0 rhombus in the bottom right corner)
// - Green: 8 (area 0 rhombus in the bottom middle)
//
// ```
//
// **Example 3:**
//
// ```
// Input: grid = [[7,7,7]]
// Output: [7]
// Explanation: All three possible rhombus sums are the same, so return [7].
//
// ```
//
// **Constraints:**
//
// - `m == grid.length`
// - `n == grid[i].length`
// - `1 <= m, n <= 50`
// - `1 <= grid[i][j] <= 105`

// This problem is so bad, bored, and has a stupid edges case, such that I don't
// even want to beat the runtime percentile.

func getBiggestThree(grid [][]int) []int {
	top := make([]int, 0, 3)
	m := len(grid)
	n := len(grid[0])

	collect := func(v int) {
		if len(top) < 3 {
			top = append(top, v)
			sort.Sort(sort.Reverse(sort.IntSlice(top)))
			return
		}

		for _, x := range top {
			if x == v {
				return
			}
		}

		i := 3
		for ; i > 0 && v > top[i-1]; i-- {
		}

		if i == 3 {
			return
		}

		copy(top[i+1:], top[i:len(top)-1])
		top[i] = v
	}

	within := func(y, x, side int) bool {
		return 0 <= y-side && y-side < m &&
			0 <= y+side && y+side < m &&
			0 <= x-side && x-side < n &&
			0 <= x+side && x+side < n
	}

	calc := func(cy, cx int) {
		collect(grid[cy][cx])

		for side := 1; within(cy, cx, side); side++ {
			y, x := cy-side, cx
			v := 0
			// top -> right
			for y < cy {
				v += grid[y][x]
				y++
				x++
			}

			for y < cy+side {
				v += grid[y][x]
				y++
				x--
			}

			for y > cy {
				v += grid[y][x]
				y--
				x--
			}

			for y > cy-side {
				v += grid[y][x]
				y--
				x++
			}
			collect(v)
		}
	}

	for y, row := range grid {
		for x := range row {
			calc(y, x)
		}
	}

	var uniq []int
	for _, v := range top {
		if len(uniq) == 0 || uniq[len(uniq)-1] != v {
			uniq = append(uniq, v)
		}
	}

	return uniq
}
