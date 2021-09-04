package main

// Pacific Atlantic Water Flow
//
// Medium
//
// https://leetcode.com/problems/pacific-atlantic-water-flow/
//
// There is an `m x n` rectangular island that borders both the **Pacific
// Ocean** and **Atlantic Ocean**. The **Pacific Ocean** touches the island's
// left and top edges, and the **Atlantic Ocean** touches the island's right and
// bottom edges.
//
// The island is partitioned into a grid of square cells. You are given an `m x
// n` integer matrix `heights` where `heights[r][c]` represents the **height
// above sea level** of the cell at coordinate `(r, c)`.
//
// The island receives a lot of rain, and the rain water can flow to neighboring
// cells directly north, south, east, and west if the neighboring cell's height
// is **less than or equal to** the current cell's height. Water can flow from
// any cell adjacent to an ocean into the ocean.
//
// Return _a **2D list** of grid coordinates_ `result` _where_ `result[i] = [ri,
// ci]` _denotes that rain water can flow from cell_`(ri, ci)` _to **both** the
// Pacific and Atlantic oceans_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/06/08/waterflow-grid.jpg)
//
// ```
// Input: heights =
// [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
// ```
//
// **Example 2:**
//
// ```
// Input: heights = [[2,1],[1,2]]
// Output: [[0,0],[0,1],[1,0],[1,1]]
//
// ```
//
// **Constraints:**
//
// - `m == heights.length`
// - `n == heights[r].length`
// - `1 <= m, n <= 200`
// - `0 <= heights[r][c] <= 105`
func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])

	red := make([][]int, m)
	blue := make([][]int, m)
	for i := range red {
		red[i] = make([]int, n)
		blue[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		u := m - 1 - y

		red[y][0] = 1
		blue[u][n-1] = 1

		for x := 1; x < n; x++ {
			v := n - 1 - x

			h := heights[y][x]
			// top
			if y == 0 ||
				// flow up
				(red[y-1][x] == 1 && heights[y-1][x] <= h) ||
				// flow left
				x > 0 && red[y][x-1] == 1 && h >= heights[y][x-1] {
				red[y][x] = 1
			}

			if red[y][x] == 1 {
				if y > 0 && red[y-1][x] == 0 && heights[y-1][x] >= h {
					red[y-1][x] = 1
				}
				if x > 0 && red[y][x-1] == 0 && heights[y][x-1] >= h {
					red[y][x-1] = 1
				}
			}

			h = heights[u][v]
			// bot
			if u == m-1 ||
				// flow down
				(blue[u+1][v] == 1 && heights[u+1][v] <= h) ||
				// flow left
				v < n-1 && blue[u][v+1] == 1 && h >= heights[u][v+1] {
				blue[u][v] = 1
			}

			if blue[u][v] == 1 {
				if u < m-1 && blue[u+1][v] == 0 && heights[u+1][v] >= h {
					blue[u+1][v] = 1
				}
				if v < n-1 && blue[u][v+1] == 0 && heights[u][v+1] >= h {
					blue[u][v+1] = 1
				}
			}
		}
	}

	printMat(red)
	printMat(blue)

	result := make([][]int, 0, m+n)
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if red[y][x] == 1 && blue[y][x] == 1 {
				result = append(result, []int{y, x})
			}
		}
	}

	return result
}

// TODO (tai): bugs, can't solve it this way, try DFS
