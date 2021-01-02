package main

import "sort"

/*
 * @lc app=leetcode id=870 lang=golang
 *
 * [870] Advantage Shuffle
 *
 * https://leetcode.com/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (43.67%)
 * Total Accepted:    24.3K
 * Total Submissions: 52.4K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * Given two arrays A and B of equal size, the advantage of A with respect to B
 * is the number of indices i for which A[i] > B[i].
 *
 * Return any permutation of A that maximizes its advantage with respect to
 * B.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [2,7,11,15], B = [1,10,4,11]
 * Output: [2,11,7,15]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [12,24,8,32], B = [13,25,32,11]
 * Output: [24,32,8,12]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length = B.length <= 10000
 * 0 <= A[i] <= 10^9
 * 0 <= B[i] <= 10^9
 *
 *
 *
 *
 */

// TODO (tai): slow, 30%
func advantageCount(a []int, b []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	sort.Ints(a)
	m := make(map[int][]int)
	for i, v := range b {
		m[v] = append(m[v], i)
	}
	sort.Ints(b)

	if b[n-1] < a[0] || b[0] > a[n-1] {
		return a
	}

	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	i := sort.SearchInts(a, b[0])
	j := 0
	for i < n && j < n {
		if a[i] == b[j] {
			i++
			continue
		}

		if a[i] < b[j] {
			i++
			continue
		}

		y := b[j]
		loc := m[y]
		res[loc[0]] = a[i]
		a[i] = -1
		m[y] = loc[1:]
		i++
		j++
	}

	for i, j := 0, 0; i < n && j < n; {
		for i < n && res[i] != -1 {
			i++
		}

		for j < n && a[j] == -1 {
			j++
		}

		if i < n && j < n {
			res[i] = a[j]
			i++
			j++
		}
	}

	return res
}
