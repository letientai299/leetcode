package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 *
 * https://leetcode.com/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Medium (33.31%)
 * Total Accepted:    187.8K
 * Total Submissions: 500.4K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * Given a list of airline tickets represented by pairs of departure and
 * arrival airports [from, to], reconstruct the itinerary in order. All of the
 * tickets belong to a man who departs from JFK. Thus, the itinerary must begin
 * with JFK.
 *
 * Note:
 *
 *
 * If there are multiple valid itineraries, you should return the itinerary
 * that has the smallest lexical order when read as a single string. For
 * example, the itinerary ["JFK", "LGA"] has a smaller lexical order than
 * ["JFK", "LGB"].
 * All airports are represented by three capital letters (IATA code).
 * You may assume all tickets form at least one valid itinerary.
 * One must use all the tickets once and only once.
 *
 *
 * Example 1:
 *
 *
 * Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
 * Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
 * Explanation: Another possible reconstruction is
 * ["JFK","SFO","ATL","JFK","ATL","SFO"].
 * But it is larger in lexical order.
 *
 *
 */
// TODO (tai): not done yet
func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	for _, ss := range tickets {
		m[ss[0]] = append(m[ss[0]], ss[1])
	}

	for k, v := range m {
		sort.Strings(v)
		m[k] = v
	}

	var res []string
	var suffix []string

	for len(m) != 0 {
		src := "JFK"
		res = append(res, src)
		dst := m[src]

		for len(dst) != 0 {
			next := dst[len(dst)-1]
			if len(dst) == 1 {
				delete(m, src)
			} else {
				m[src] = dst[:len(dst)-1]
			}
			res = append(res, next)
			src = next
			dst = m[src]
		}

		if len(m) != 0 {
			suffix = append(res[1:], suffix...)
			res = make([]string, 0, len(suffix))
		}
	}
	res = append(res, suffix...)
	return res
}
