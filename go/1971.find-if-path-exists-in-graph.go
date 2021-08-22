package main

// Find if Path Exists in Graph
//
// Easy
//
// https://leetcode.com/problems/find-if-path-exists-in-graph/
//
// There is a **bi-directional** graph with `n` vertices, where each vertex is
// labeled from `0` to `n - 1` ( **inclusive**). The edges in the graph are
// represented as a 2D integer array `edges`, where each `edges[i] = [ui, vi]`
// denotes a bi-directional edge between vertex `ui` and vertex `vi`. Every
// vertex pair is connected by **at most one** edge, and no vertex has an edge
// to itself.
//
// You want to determine if there is a **valid path** that exists from vertex
// `start` to vertex `end`.
//
// Given `edges` and the integers `n`, `start`, and `end`, return `true` _if
// there is a **valid path** from_ `start` _to_ `end` _, or_ `false` _otherwise_
// _._
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/14/validpath-ex1.png)
//
// ```
// Input: n = 3, edges = [[0,1],[1,2],[2,0]], start = 0, end = 2
// Output: true
// Explanation: There are two paths from vertex 0 to vertex 2:
// - 0 → 1 → 2
// - 0 → 2
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/14/validpath-ex2.png)
//
// ```
// Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], start = 0, end = 5
// Output: false
// Explanation: There is no path from vertex 0 to vertex 5.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 2 * 105`
// - `0 <= edges.length <= 2 * 105`
// - `edges[i].length == 2`
// - `1 <= ui, vi <= n - 1`
// - `ui != vi`
// - `1 <= start, end <= n - 1`
// - There are no duplicate edges.
// - There are no self edges.
func validPath(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}

	g := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		g = append(g, nil)
	}

	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	reachable := make(map[int]bool)
	q := []int{start}
	for len(q) != 0 {
		var next []int

		for _, x := range q {
			for _, y := range g[x] {
				if !reachable[y] {
					reachable[y] = true
					next = append(next, y)
					if y == end {
						return true
					}
				}
			}
		}

		q = next
	}

	return false
}
