package main

import "sort"

/*
 * @lc app=leetcode id=881 lang=golang
 *
 * [881] Boats to Save People
 *
 * https://leetcode.com/problems/boats-to-save-people/description/
 *
 * algorithms
 * Medium (45.00%)
 * Total Accepted:    42.3K
 * Total Submissions: 89K
 * Testcase Example:  '[1,2]\n3'
 *
 * The i-th person has weight people[i], and each boat can carry a maximum
 * weight of limit.
 *
 * Each boat carries at most 2 people at the same time, provided the sum of the
 * weight of those people is at most limit.
 *
 * Return the minimum number of boats to carry every given person.  (It is
 * guaranteed each person can be carried by a boat.)
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: people = [1,2], limit = 3
 * Output: 1
 * Explanation: 1 boat (1, 2)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: people = [3,2,2,1], limit = 3
 * Output: 3
 * Explanation: 3 boats (1, 2), (2) and (3)
 *
 *
 *
 * Example 3:
 *
 *
 * Input: people = [3,5,3,4], limit = 5
 * Output: 4
 * Explanation: 4 boats (3), (3), (4), (5)
 *
 * Note:
 *
 *
 * 1 <= people.length <= 50000
 * 1 <= people[i] <= limit <= 30000
 *
 *
 *
 *
 *
 */
func numRescueBoats(people []int, limit int) int {
	n := len(people)
	if n == 1 {
		return 1
	}

	sort.Ints(people)
	b := 0
	for i, j := 0, n-1; i <= j; {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}
		b++
	}
	return b
}
