package main

/*
 * @lc app=leetcode id=1128 lang=golang
 *
 * [1128] Number of Equivalent Domino Pairs
 *
 * https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/
 *
 * algorithms
 * Easy (46.20%)
 * Total Accepted:    28.2K
 * Total Submissions: 60.1K
 * Testcase Example:  '[[1,2],[2,1],[3,4],[5,6]]'
 *
 * Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j]
 * = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that
 * is, one domino can be rotated to be equal to another domino.
 *
 * Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length,
 * and dominoes[i] is equivalent to dominoes[j].
 *
 *
 * Example 1:
 * Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 1 <= dominoes.length <= 40000
 * 1 <= dominoes[i][j] <= 9
 *
 */
func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[int]int)
	for _, d := range dominoes {
		x, y := d[0], d[1]
		if x > y {
			x, y = y, x
		}
		m[x*10+y]++
	}

	s := 0
	for _, v := range m {
		s += v * (v - 1) / 2
	}
	return s
}
