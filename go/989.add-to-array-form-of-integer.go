package main

/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 *
 * https://leetcode.com/problems/add-to-array-form-of-integer/description/
 *
 * algorithms
 * Easy (44.02%)
 * Total Accepted:    52.8K
 * Total Submissions: 118.7K
 * Testcase Example:  '[1,2,0,0]\n34'
 *
 * For a non-negative integer X, the array-form of X is an array of its digits
 * in left to right order.  For example, if X = 1231, then the array form is
 * [1,2,3,1].
 *
 * Given the array-form A of a non-negative integer X, return the array-form of
 * the integer X+K.
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
 * Input: A = [1,2,0,0], K = 34
 * Output: [1,2,3,4]
 * Explanation: 1200 + 34 = 1234
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [2,7,4], K = 181
 * Output: [4,5,5]
 * Explanation: 274 + 181 = 455
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2,1,5], K = 806
 * Output: [1,0,2,1]
 * Explanation: 215 + 806 = 1021
 *
 *
 *
 * Example 4:
 *
 *
 * Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
 * Output: [1,0,0,0,0,0,0,0,0,0,0]
 * Explanation: 9999999999 + 1 = 10000000000
 *
 *
 *
 *
 * Note：
 *
 *
 * 1 <= A.length <= 10000
 * 0 <= A[i] <= 9
 * 0 <= K <= 10000
 * If A.length > 1, then A[0] != 0
 *
 *
 *
 *
 *
 */
func addToArrayForm(a []int, k int) []int {
	i := len(a) - 1
	for ; i >= 0 && k > 0; i-- {
		a[i] += k % 10
		rem := a[i] / 10
		a[i] %= 10
		k /= 10
		k += rem
	}

	var prefix []int
	for k > 0 {
		prefix = append(prefix, k%10)
		k /= 10
	}

	for i := 0; i < len(prefix)/2; i++ {
		prefix[i], prefix[len(prefix)-1-i] = prefix[len(prefix)-1-i], prefix[i]
	}

	return append(prefix, a...)
}
