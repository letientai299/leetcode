package main

// Number of Operations to Make Network Connected
//
// Medium
//
// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
//
// There are `n` computers numbered from `0` to `n-1` connected by ethernet
// cables `connections` forming a network where `connections[i] = [a,
// b]` represents a connection between computers `a` and `b`. Any computer can
// reach any other computer directly or indirectly through the network.
//
// Given an initial computer network `connections`. You can extract certain
// cables between two directly connected computers, and place them between any
// pair of disconnected computers to make them directly connected. Return the
// _minimum number of times_ you need to do this in order to make all the
// computers connected. If it's not possible, return -1.
//
// **Example 1:**
//
// **![](https://assets.leetcode.com/uploads/2020/01/02/sample_1_1677.png)**
//
// ```
// Input: n = 4, connections = [[0,1],[0,2],[1,2]]
// Output: 1
// Explanation: Remove cable between computer 1 and 2 and place between
// computers 1 and 3.
//
// ```
//
// **Example 2:**
//
// **![](https://assets.leetcode.com/uploads/2020/01/02/sample_2_1677.png)**
//
// ```
// Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
// Output: 2
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
// Output: -1
// Explanation: There are not enough cables.
//
// ```
//
// **Example 4:**
//
// ```
// Input: n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 10^5`
// - `1 <= connections.length <= min(n*(n-1)/2, 10^5)`
// - `connections[i].length == 2`
// - `0 <= connections[i][0], connections[i][1] < n`
// - `connections[i][0] != connections[i][1]`
// - There are no repeated connections.
// - No two computers are connected by more than one cable.
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1 // not enough edges to build a tree from given graph
	}

	type set map[int]struct{}
	m := make(map[int]*set)

	// count number of disjoint graphs
	for _, c := range connections {
		x := c[0]
		y := c[1]
		if m[x] == nil && m[y] == nil {
			s := set{}
			s[x] = struct{}{}
			s[y] = struct{}{}
			m[x] = &s
			m[y] = &s
			continue
		}

		if m[x] != nil && m[y] == nil {
			m[y] = m[x]
			(*m[x])[y] = struct{}{}
			continue
		}

		if m[x] == nil && m[y] != nil {
			m[x] = m[y]
			(*m[y])[x] = struct{}{}
			continue
		}

		if m[x] != m[y] {
			a, b := m[x], m[y]
			for k, v := range *b {
				(*a)[k] = v
			}

			if len(*m[x]) < len(*m[y]) {
				x, y = y, x
			}

			for node := range *m[y] {
				m[node] = m[x]
			}
		}
	}

	count := make(map[*set]struct{})
	for _, v := range m {
		count[v] = struct{}{}
	}

	return n - len(m) + len(count) - 1
}

// TODO (tai): 1200ms, 14.29%! solve it again, must have a nicer and faster solution
