package main

/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 *
 * https://leetcode.com/problems/evaluate-division/description/
 *
 * algorithms
 * Medium (49.33%)
 * Total Accepted:    165.6K
 * Total Submissions: 307.6K
 * Testcase Example:  '[["a","b"],["b","c"]]\n' +
  '[2.0,3.0]\n' +
  '[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'
 *
 * You are given an array of variable pairs equations and an array of real
 * numbers values, where equations[i] = [Ai, Bi] and values[i] represent the
 * equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a
 * single variable.
 *
 * You are also given some queries, where queries[j] = [Cj, Dj] represents the
 * j^th query where you must find the answer for Cj / Dj = ?.
 *
 * Return the answers to all queries. If a single answer cannot be determined,
 * return -1.0.
 *
 * Note: The input is always valid. You may assume that evaluating the queries
 * will not result in division by zero and that there is no contradiction.
 *
 *
 * Example 1:
 *
 *
 * Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries =
 * [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
 * Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
 * Explanation:
 * Given: a / b = 2.0, b / c = 3.0
 * queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
 * return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
 *
 *
 * Example 2:
 *
 *
 * Input: equations = [["a","b"],["b","c"],["bc","cd"]], values =
 * [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
 * Output: [3.75000,0.40000,5.00000,0.20000]
 *
 *
 * Example 3:
 *
 *
 * Input: equations = [["a","b"]], values = [0.5], queries =
 * [["a","b"],["b","a"],["a","c"],["x","y"]]
 * Output: [0.50000,2.00000,-1.00000,-1.00000]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= equations.length <= 20
 * equations[i].length == 2
 * 1 <= Ai.length, Bi.length <= 5
 * values.length == equations.length
 * 0.0 < values[i] <= 20.0
 * 1 <= queries.length <= 20
 * queries[i].length == 2
 * 1 <= Cj.length, Dj.length <= 5
 * Ai, Bi, Cj, Dj consist of lower case English letters and digits.
 *
 *
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]map[string]float64)
	for i, e := range equations {
		v := values[i]
		a, b := e[0], e[1]
		if m[a] == nil {
			m[a] = make(map[string]float64)
		}
		m[a][b] = v
		if m[b] == nil {
			m[b] = make(map[string]float64)
		}
		m[b][a] = 1 / v
	}

	var find func(a, b string, visited map[string]struct{}) float64

	find = func(a, b string, visited map[string]struct{}) float64 {
		for x, v := range m[a] {
			if _, ok := visited[x]; ok {
				continue
			}

			if x == b {
				return v
			}

			visited[x] = struct{}{}
			if u := find(x, b, visited); u != -1 {
				return u * v
			}
		}
		return -1
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]

		if m[a] == nil || m[b] == nil {
			res[i] = -1
			continue
		}

		res[i] = find(a, b, map[string]struct{}{})
	}

	return res
}
