package main

/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (71.86%)
 * Total Accepted:    332.3K
 * Total Submissions: 460.9K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * Given an array of integers A sorted in non-decreasing order, return an array
 * of the squares of each number, also in sorted non-decreasing order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A is sorted in non-decreasing order.
 *
 *
 *
 */
func sortedSquares(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return []int{a[0] * a[0]}
	}

	i := 1
	for ; i < n; i++ {
		if a[i-1]*a[i] <= 0 {
			break
		}
	}

	res := make([]int, 0, n)
	if i == n {
		if a[0] < 0 {
			for i := n - 1; i >= 0; i-- {
				res = append(res, a[i]*a[i])
			}
		} else {
			for i := 0; i < n; i++ {
				res = append(res, a[i]*a[i])
			}
		}
		return res
	}

	j := i - 1
	for ; j >= 0 && i < n; {
		if -a[j] == a[i] {
			res = append(res, a[j]*a[j], a[i]*a[i])
			j--
			i++
		} else if -a[j] < a[i] {
			res = append(res, a[j]*a[j])
			j--
		} else {
			res = append(res, a[i]*a[i])
			i++
		}
	}

	for ; i < n; i++ {
		res = append(res, a[i]*a[i])
	}
	for ; j >= 0; j-- {
		res = append(res, a[j]*a[j])
	}

	return res
}
