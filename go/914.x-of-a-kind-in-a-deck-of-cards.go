package main

/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 *
 * https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/description/
 *
 * algorithms
 * Easy (34.01%)
 * Total Accepted:    53.7K
 * Total Submissions: 155.3K
 * Testcase Example:  '[1,2,3,4,4,3,2,1]'
 *
 * In a deck of cards, each card has an integer written on it.
 *
 * Return true if and only if you can choose X >= 2 such that it is possible to
 * split the entire deck into 1 or more groups of cards, where:
 *
 *
 * Each group has exactly X cards.
 * All the cards in each group have the same integer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: deck = [1,2,3,4,4,3,2,1]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
 *
 *
 * Example 2:
 *
 *
 * Input: deck = [1,1,1,2,2,2,3,3]
 * Output: false´
 * Explanation: No possible partition.
 *
 *
 * Example 3:
 *
 *
 * Input: deck = [1]
 * Output: false
 * Explanation: No possible partition.
 *
 *
 * Example 4:
 *
 *
 * Input: deck = [1,1]
 * Output: true
 * Explanation: Possible partition [1,1].
 *
 *
 * Example 5:
 *
 *
 * Input: deck = [1,1,2,2,2,2]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[2,2].
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= deck.length <= 10^4
 * 0 <= deck[i] < 10^4
 *
 *
 */
func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, x := range deck {
		m[x]++
	}

	var gcd func(a, b int) int

	gcd = func(a, b int) int {
		if a == 1 {
			return 1
		}

		if a == 0 {
			return b
		}

		if a > b {
			return gcd(b, a)
		}

		return gcd(b%a, a)
	}

	n := len(deck)
	for _, v := range m {
		n = gcd(v, n)
		if n == 1 {
			return false
		}
	}

	return n > 1
}
