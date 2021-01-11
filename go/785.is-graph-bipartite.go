package main

/*
 * @lc app=leetcode id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 *
 * https://leetcode.com/problems/is-graph-bipartite/description/
 *
 * algorithms
 * Medium (45.39%)
 * Total Accepted:    154.8K
 * Total Submissions: 321.4K
 * Testcase Example:  '[[1,3],[0,2],[1,3],[0,2]]'
 *
 * Given an undirected graph, return true if and only if it is bipartite.
 *
 * Recall that a graph is bipartite if we can split its set of nodes into two
 * independent subsets A and B, such that every edge in the graph has one node
 * in A and another node in B.
 *
 * The graph is given in the following form: graph[i] is a list of indexes j
 * for which the edge between nodes i and j exists.  Each node is an integer
 * between 0 and graph.length - 1.  There are no self edges or parallel edges:
 * graph[i] does not contain i, and it doesn't contain any element twice.
 *
 *
 * Example 1:
 *
 *
 * Input: graph = [[1,3],[0,2],[1,3],[0,2]]
 * Output: true
 * Explanation: We can divide the vertices into two groups: {0, 2} and {1,
 * 3}.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
 * Output: false
 * Explanation: We cannot find a way to divide the set of nodes into two
 * independent subsets.
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= graph.length <= 100
 * 0 <= graph[i].length < 100
 * 0 <= graph[i][j] <= graph.length - 1
 * graph[i][j] != i
 * All the values of graph[i] are unique.
 * The graph is guaranteed to be undirected.
 *
 *
 */

func isBipartite(graph [][]int) bool {
	visited := make(map[int]bool)

	var visit func(e int, left bool) bool
	visit = func(e int, left bool) bool {
		if where, ok := visited[e]; ok {
			return where == left
		}

		visited[e] = left
		for _, v := range graph[e] {
			if !visit(v, !left) {
				return false
			}
		}
		return true
	}

	for i, v := range graph {
		if _, ok := visited[i]; ok || len(v) == 0 {
			continue
		}

		if !visit(i, true) {
			return false
		}
	}

	return true
}
