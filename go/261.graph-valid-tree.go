package main

// Graph Valid Tree
//
// # Medium
//
// https://leetcode.com/problems/graph-valid-tree
func validTree(n int, edges [][]int) bool {
	m := make([][]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}

	// start from 0, check if we can discover all other node without revisit any one
	seen := make(map[int]bool)

	var dfs func(x int, pre int) bool
	dfs = func(x int, pre int) bool {
		if seen[x] {
			return false
		}

		seen[x] = true
		for _, y := range m[x] {
			if y != pre && !dfs(y, x) {
				return false
			}
		}

		return true
	}

	return dfs(0, -1) && len(seen) == n // all nodes are connected.
}
