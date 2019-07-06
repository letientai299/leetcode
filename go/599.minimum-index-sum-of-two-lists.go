package main

import (
	"math"
)

/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 *
 * https://leetcode.com/problems/minimum-index-sum-of-two-lists/description/
 *
 * algorithms
 * Easy (48.19%)
 * Total Accepted:    59.4K
 * Total Submissions: 123.1K
 * Testcase Example:  '["Shogun","Tapioca Express","Burger King","KFC"]\n["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]'
 *
 *
 * Suppose Andy and Doris want to choose a restaurant for dinner, and they both
 * have a list of favorite restaurants represented by strings.
 *
 *
 * You need to help them find out their common interest with the least list
 * index sum. If there is a choice tie between answers, output all of them with
 * no order requirement. You could assume there always exists an answer.
 *
 *
 *
 * Example 1:
 *
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
 * Output: ["Shogun"]
 * Explanation: The only restaurant they both like is "Shogun".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["KFC", "Shogun", "Burger King"]
 * Output: ["Shogun"]
 * Explanation: The restaurant they both like and have the least index sum is
 * "Shogun" with index sum 1 (0+1).
 *
 *
 *
 *
 * Note:
 *
 * The length of both lists will be in the range of [1, 1000].
 * The length of strings in both lists will be in the range of [1, 30].
 * The index is starting from 0 to the list length minus 1.
 * No duplicates in both lists.
 *
 *
 */
func findRestaurant(list1 []string, list2 []string) []string {
	map1 := make(map[string]int, len(list1))
	for i, s := range list1 {
		map1[s] = i
	}

	var matches []string
	minIndexSum := math.MaxInt32
	for idx2, s := range list2 {
		idx1, exits := map1[s]
		if !exits {
			continue
		}

		idxSum := idx1 + idx2

		if len(matches) == 0 {
			minIndexSum = idxSum
			matches = append(matches, s)
			continue
		}

		if minIndexSum < idxSum {
			continue
		}

		if minIndexSum == idxSum {
			matches = append(matches, s)
			continue
		}

		minIndexSum = idxSum
		matches = []string{s}
	}

	return matches
}
