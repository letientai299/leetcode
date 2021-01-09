package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=846 lang=golang
 *
 * [846] Hand of Straights
 *
 * https://leetcode.com/problems/hand-of-straights/description/
 *
 * algorithms
 * Medium (50.65%)
 * Total Accepted:    62.6K
 * Total Submissions: 113.5K
 * Testcase Example:  '[1,2,3,6,2,3,4,7,8]\n3'
 *
 * Alice has a hand of cards, given as an array of integers.
 *
 * Now she wants to rearrange the cards into groups so that each group is size
 * W, and consists of W consecutive cards.
 *
 * Return true if and only if she can.
 *
 * Note: This question is the same as 1296:
 * https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
 *
 *
 * Example 1:
 *
 *
 * Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
 * Output: true
 * Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
 *
 *
 * Example 2:
 *
 *
 * Input: hand = [1,2,3,4,5], W = 4
 * Output: false
 * Explanation: Alice's hand can't be rearranged into groups of 4.
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= hand.length <= 10000
 * 0 <= hand[i] <= 10^9
 * 1 <= W <= hand.length
 *
 *
 */
func isNStraightHand(hand []int, k int) bool {
	if k == 1 {
		return true
	}

	if len(hand)%k != 0 {
		return false
	}

	m := make(map[int]int)
	var uniq []int
	for _, v := range hand {
		if m[v] == 0 {
			uniq = append(uniq, v)
		}
		m[v]++
	}

	if len(uniq) < k {
		return false
	}

	sort.Ints(uniq)

	for i := 0; i < len(uniq); i++ {
		v := uniq[i]
		if m[v] == 0 {
			continue
		}
		p := m[v]
		m[v] = 0
		for j := 1; j < k; j++ {
			if m[v+j] < p {
				return false
			}

			m[v+j] -= p
			if m[v+j] == 0 {
				if m[v+j-1] > 0 {
					return false
				}
				i++
			}
		}
	}

	return true
}
