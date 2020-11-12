package main

import "sort"

/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 *
 * https://leetcode.com/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (57.28%)
 * Total Accepted:    38.2K
 * Total Submissions: 65.5K
 * Testcase Example:  '[2,1,2]'
 *
 * Given an array A of positive lengths, return the largest perimeter of a
 * triangle with non-zero area, formed from 3 of these lengths.
 *
 * If it is impossible to form any triangle of non-zero area, return 0.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1,2]
 * Output: 5
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,1]
 * Output: 0
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [3,2,3,4]
 * Output: 10
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [3,6,2,3]
 * Output: 8
 *
 *
 *
 *
 * Note:
 *
 *
 * 3 <= A.length <= 10000
 * 1 <= A[i] <= 10^6
 *
 *
 *
 *
 *
 */
func largestPerimeter(arr []int) int {
	sort.Ints(arr)

	tri := func(a, b, c int) bool {
		return (a+b) > c && (b+c) > a && (c+a) > b
	}

	for i := len(arr) - 1; i >= 2; i-- {
		triangle := true
		for j := i - 1; j >= 1 && triangle; j-- {
			for k := j - 1; k >= 0 && triangle; k-- {
				triangle = tri(arr[i], arr[j], arr[k])
				if triangle {
					return arr[i] + arr[j] + arr[k]
				}
			}
		}
	}

	return 0
}
