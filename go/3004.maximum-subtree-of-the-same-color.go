package main

// Maximum Subtree of the Same Color
//
// # Medium
//
// https://leetcode.com/problems/maximum-subtree-of-the-same-color
func maximumSubtreeSize(edges [][]int, colors []int) int {
	n := len(edges) + 1
	nodes := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		nodes[a] = append(nodes[a], b)
		nodes[b] = append(nodes[b], a)
	}

	seen := make([]bool, n)
	var bfs func(root int) (int, bool)
	bfs = func(root int) (int, bool) {
		seen[root] = true

		r := 1
		same := true
		best := 0
		for _, a := range nodes[root] {
			if seen[a] {
				continue
			}

			x, ok := bfs(a)
			same = same && colors[a] == colors[root] && ok
			best = max(best, x)
			r += x
		}

		if same {
			return r, same
		}
		return best, same
	}

	r, _ := bfs(0)
	return r
}
