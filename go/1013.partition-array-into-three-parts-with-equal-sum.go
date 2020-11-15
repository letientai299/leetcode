package main

/*
 * @lc app=leetcode id=1013 lang=golang
 *
 * [1013] Partition Array Into Three Parts With Equal Sum
 *
 * https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/description/
 *
 * algorithms
 * Easy (56.67%)
 * Total Accepted:    39.6K
 * Total Submissions: 78.9K
 * Testcase Example:  '[0,2,1,-6,6,-7,9,1,2,0,1]'
 *
 * Given an array A of integers, return true if and only if we can partition
 * the array into three non-empty parts with equal sums.
 *
 * Formally, we can partition the array if we can find indexes i+1 < j with
 * (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1]
 * + ... + A[A.length - 1])
 *
 *
 * Example 1:
 *
 *
 * Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
 * Output: true
 * Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
 *
 *
 * Example 2:
 *
 *
 * Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: A = [3,3,6,5,-2,2,5,1,-9,4]
 * Output: true
 * Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= A.length <= 50000
 * -10^4 <= A[i] <= 10^4
 *
 *
 */

func canThreePartsEqualSum(a []int) bool {
	s := 0
	sum := make([]int, len(a))
	for i, x := range a {
		if i == 0 {
			sum[0] = a[0]
		} else {
			sum[i] = sum[i-1] + a[i]
		}
		s += x
	}

	if s%3 != 0 {
		return false
	}

	p1, p2 := false, false
	i1 := -1
	for i, x := range sum {
		if x == s/3 && i1 == -1 {
			p1 = true
			i1 = i
		}
		if x == 2*s/3 && i > i1 && i1 >= 0 && i < len(a)-1 {
			p2 = true
			break
		}
	}
	return p1 && p2
}
