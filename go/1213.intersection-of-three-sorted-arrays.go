package main

/*
 * @lc app=leetcode id=1213 lang=golang
 *
 * [1213] Intersection of Three Sorted Arrays
 *
 * https://leetcode.com/problems/intersection-of-three-sorted-arrays/description/
 *
 * algorithms
 * Easy (77.07%)
 * Total Accepted:    32.5K
 * Total Submissions: 41.1K
 * Testcase Example:  '[1,2,3,4,5]\n[1,2,5,7,9]\n[1,3,4,5,8]'
 *
 * Given three integer arrays arr1, arr2 and arr3 sorted in strictly increasing
 * order, return a sorted array of only the integers that appeared in all three
 * arrays.
 *
 *
 * Example 1:
 *
 *
 * Input: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
 * Output: [1,5]
 * Explanation: Only 1 and 5 appeared in the three arrays.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr1.length, arr2.length, arr3.length <= 1000
 * 1 <= arr1[i], arr2[i], arr3[i] <= 2000
 *
 *
 */
func arraysIntersection(a []int, b []int, c []int) []int {
	x, y, z := 0, 0, 0
	m, n, p := len(a), len(b), len(c)
	var res []int
	for x < m && y < n && z < p {
		if a[x] == b[y] && b[y] == c[z] {
			res = append(res, a[x])
			x++
			y++
			z++
			continue
		}

		if a[x] < b[y] || a[x] < c[z] {
			x++
			continue
		}

		if b[y] < c[z] || b[y] < a[x] {
			y++
			continue
		}

		if c[z] < a[x] || c[z] < b[y] {
			z++
			continue
		}
	}

	return res
}
