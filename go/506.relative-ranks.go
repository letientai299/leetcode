package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 *
 * https://leetcode.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (48.16%)
 * Total Accepted:    41.5K
 * Total Submissions: 86.2K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 *
 * Given scores of N athletes, find their relative ranks and the people with
 * the top three highest scores, who will be awarded medals: "Gold Medal",
 * "Silver Medal" and "Bronze Medal".
 *
 * Example 1:
 *
 * Input: [5, 4, 3, 2, 1]
 * Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * Explanation: The first three athletes got the top three highest scores, so
 * they got "Gold Medal", "Silver Medal" and "Bronze Medal". For the left two
 * athletes, you just need to output their relative ranks according to their
 * scores.
 *
 *
 *
 * Note:
 *
 * N is a positive integer and won't exceed 10,000.
 * All the scores of athletes are guaranteed to be unique.
 *
 *
 *
 */
func findRelativeRanks(nums []int) []string {
	n := len(nums)
	if n <= 0 {
		return nil
	}

	indexes := make(map[int]int)
	for index, val := range nums {
		indexes[val] = index
	}

	sort.Ints(nums)
	res := make([]string, n)

	// process first 3 ranks, since they're special
	res[indexes[nums[n-1]]] = "Gold Medal"

	if n < 2 {
		return res
	}
	res[indexes[nums[n-2]]] = "Silver Medal"

	if n < 3 {
		return res
	}
	res[indexes[nums[n-3]]] = "Bronze Medal"

	// note that at this point the array nums is sorted
	for reverseRank := n - 4; reverseRank >= 0; reverseRank-- {
		originalIndex := indexes[nums[reverseRank]]
		rank := n - reverseRank
		res[originalIndex] = strconv.Itoa(rank)
	}

	return res
}
