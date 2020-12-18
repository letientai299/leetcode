package main

import (
	"math"
)

/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 *
 * https://leetcode.com/problems/minimum-height-trees/description/
 *
 * algorithms
 * Medium (31.11%)
 * Total Accepted:    123.9K
 * Total Submissions: 360.7K
 * Testcase Example:  '4\n[[1,0],[1,2],[1,3]]'
 *
 * A tree is an undirected graph in which any two vertices are connected by
 * exactly one path. In other words, any connected graph without simple cycles
 * is a tree.
 *
 * Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1
 * edges where edges[i] = [ai, bi] indicates that there is an undirected edge
 * between the two nodes ai and bi in the tree, you can choose any node of the
 * tree as the root. When you select a node x as the root, the result tree has
 * height h. Among all possible rooted trees, those with minimum height (i.e.
 * min(h))  are called minimum height trees (MHTs).
 *
 * Return a list of all MHTs' root labels. You can return the answer in any
 * order.
 *
 * The height of a rooted tree is the number of edges on the longest downward
 * path between the root and a leaf.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4, edges = [[1,0],[1,2],[1,3]]
 * Output: [1]
 * Explanation: As shown, the height of the tree is 1 when the root is the node
 * with label 1 which is the only MHT.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
 * Output: [3,4]
 *
 *
 * Example 3:
 *
 *
 * Input: n = 1, edges = []
 * Output: [0]
 *
 *
 * Example 4:
 *
 *
 * Input: n = 2, edges = [[0,1]]
 * Output: [0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 2 * 10^4
 * edges.length == n - 1
 * 0 <= ai, bi < n
 * ai != bi
 * All the pairs (ai, bi) are distinct.
 * The given input is guaranteed to be a tree and there will be no repeated
 * edges.
 *
 *
 */

// TODO (tai): not done yet, time limit exceed
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	}

	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	m := make(map[int][]int)
	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
		m[e[1]] = append(m[e[1]], e[0])
	}

	var height func(x, p int) int

	height = func(x, p int) int {
		h := 1 // root node
		if len(m[x]) == 0 || (len(m[x]) == 1 && m[x][0] == p) {
			return 1
		}

		for _, v := range m[x] {
			if v == p {
				// don't go back up the tree
				continue
			}

			tmp := height(v, x)
			if tmp+1 > h {
				h = tmp + 1
			}
		}

		return h
	}

	best := math.MaxInt32
	var r []int
	for i := 0; i < n; i++ {
		if len(m[i]) < 2 {
			continue
		}

		h := height(i, -1)
		if h < best {
			best = h
			r = []int{i}
			continue
		}

		if h == best {
			r = append(r, i)
		}
	}

	return r
}
