package main

// Find Eventual Safe States
//
// Medium
//
// https://leetcode.com/problems/find-eventual-safe-states/
//
// We start at some node in a directed graph, and every turn, we walk along a
// directed edge of the graph. If we reach a terminal node (that is, it has no
// outgoing directed edges), we stop.
//
// We define a starting node to be **safe** if we must eventually walk to a
// terminal node. More specifically, there is a natural number `k`, so that we
// must have stopped at a terminal node in less than `k` steps for **any choice
// of where to walk**.
//
// Return _an array containing all the safe nodes of the graph_. The answer
// should be sorted in **ascending** order.
//
// The directed graph has `n` nodes with labels from `0` to `n - 1`, where `n`
// is the length of `graph`. The graph is given in the following form:
// `graph[i]` is a list of labels `j` such that `(i, j)` is a directed edge of
// the graph, going from node `i` to node `j`.
//
// **Example 1:**
//
// ![Illustration of
// graph](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/03/17/picture1.png)
//
// ```
// Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
// Output: [2,4,5,6]
// Explanation: The given graph is shown above.
//
// ```
//
// **Example 2:**
//
// ```
// Input: graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]
// Output: [4]
//
// ```
//
// **Constraints:**
//
// - `n == graph.length`
// - `1 <= n <= 104`
// - `0 <= graph[i].length <= n`
// - `graph[i]` is sorted in a strictly increasing order.
// - The graph may contain self-loops.
// - The number of edges in the graph will be in the range `[1, 4 * 104]`.
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	unsafes := make(map[int]bool)

	seen := make(map[int]bool)
	var walk func(cur int)

	walk = func(cur int) {
		if seen[cur] {
			for x, saw := range seen {
				if saw {
					unsafes[x] = true
				}
			}
			return
		}

		seen[cur] = true
		for _, v := range graph[cur] {
			if unsafes[v] || v == cur {
				unsafes[cur] = true
				break
			}

			walk(v)
		}
		seen[cur] = false

		if unsafes[cur] {
			for x, saw := range seen {
				if saw {
					unsafes[x] = true
				}
			}
			return
		}

	}

	for v := range graph {
		if unsafes[v] {
			continue
		}

		walk(v)
		if len(unsafes) == n {
			break
		}
	}

	var res []int
	for i := 0; i < n; i++ {
		if !unsafes[i] {
			res = append(res, i)
		}
	}

	return res
}

// TODO (tai): slow and ugly, TLE
