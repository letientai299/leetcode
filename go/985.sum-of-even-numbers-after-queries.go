package main

/*
 * @lc app=leetcode id=985 lang=golang
 *
 * [985] Sum of Even Numbers After Queries
 *
 * https://leetcode.com/problems/sum-of-even-numbers-after-queries/description/
 *
 * algorithms
 * Easy (62.35%)
 * Total Accepted:    47.3K
 * Total Submissions: 77.7K
 * Testcase Example:  '[1,2,3,4]\n[[1,0],[-3,1],[-4,0],[2,3]]'
 *
 * We have an array A of integers, and an array queries of queries.
 *
 * For the i-th query val = queries[i][0], index = queries[i][1], we add val to
 * A[index].  Then, the answer to the i-th query is the sum of the even values
 * of A.
 *
 * (Here, the given index = queries[i][1] is a 0-based index, and each query
 * permanently modifies the array A.)
 *
 * Return the answer to all queries.  Your answer array should have answer[i]
 * as the answer to the i-th query.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
 * Output: [8,6,2,4]
 * Explanation:
 * At the beginning, the array is [1,2,3,4].
 * After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values
 * is 2 + 2 + 4 = 8.
 * After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values
 * is 2 + 4 = 6.
 * After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even
 * values is -2 + 4 = 2.
 * After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values
 * is -2 + 6 = 4.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * 1 <= queries.length <= 10000
 * -10000 <= queries[i][0] <= 10000
 * 0 <= queries[i][1] < A.length
 *
 *
 */
func sumEvenAfterQueries(a []int, queries [][]int) []int {
	se := 0
	for _, x := range a {
		if x%2 == 0 {
			se += x
		}
	}
	res := make([]int, 0, len(queries))
	for _, q := range queries {
		if a[q[1]]%2 == 0 {
			se -= a[q[1]]
		}

		a[q[1]] += q[0]

		if a[q[1]]%2 == 0 {
			se += a[q[1]]
		}

		res = append(res, se)
	}
	return res
}
