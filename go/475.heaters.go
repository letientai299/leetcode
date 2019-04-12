package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=475 lang=golang
 *
 * [475] Heaters
 *
 * https://leetcode.com/problems/heaters/description/
 *
 * algorithms
 * Easy (31.46%)
 * Total Accepted:    45.6K
 * Total Submissions: 145.1K
 * Testcase Example:  '[1,2,3]\n[2]'
 *
 * Winter is coming! Your first job during the contest is to design a standard
 * heater with fixed warm radius to warm all the houses.
 *
 * Now, you are given positions of houses and heaters on a horizontal line,
 * find out minimum radius of heaters so that all houses could be covered by
 * those heaters.
 *
 * So, your input will be the positions of houses and heaters separately, and
 * your expected output will be the minimum radius standard of heaters.
 *
 * Note:
 *
 *
 * Numbers of houses and heaters you are given are non-negative and will not
 * exceed 25000.
 * Positions of houses and heaters you are given are non-negative and will not
 * exceed 10^9.
 * As long as a house is in the heaters' warm radius range, it can be
 * warmed.
 * All the heaters follow your radius standard and the warm radius will the
 * same.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3],[2]
 * Output: 1
 * Explanation: The only heater was placed in the position 2, and if we use the
 * radius 1 standard, then all the houses can be warmed.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4],[1,4]
 * Output: 1
 * Explanation: The two heater was placed in the position 1 and 4. We need to
 * use radius 1 standard, then all the houses can be warmed.
 *
 *
 *
 *
 */
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	radius := 0
	for _, h := range houses {
		loc := sort.SearchInts(heaters, h)
		left := int(math.MaxInt32)
		right := left
		if loc >= len(heaters) {
			left = h - heaters [len(heaters)-1]
		} else if heaters[loc] == h {
			continue
		} else if loc == 0 {
			right = heaters[0] - h
		} else {
			left = h - heaters[loc-1]
			right = heaters[loc] - h
		}

		if left < right && radius < left {
			radius = left
			continue
		}

		if radius < right && right < left {
			radius = right
		}
	}

	return radius
}
