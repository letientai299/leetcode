package main

/*
 * @lc app=leetcode id=877 lang=golang
 *
 * [877] Stone Game
 *
 * https://leetcode.com/problems/stone-game/description/
 *
 * algorithms
 * Medium (62.75%)
 * Total Accepted:    72K
 * Total Submissions: 109K
 * Testcase Example:  '[5,3,4,5]'
 *
 * Alex and Lee play a game with piles of stones.  There are an even number of
 * piles arranged in a row, and each pile has a positive integer number of
 * stones piles[i].
 *
 * The objective of the game is to end with the most stones.  The total number
 * of stones is odd, so there are no ties.
 *
 * Alex and Lee take turns, with Alex starting first.  Each turn, a player
 * takes the entire pile of stones from either the beginning or the end of the
 * row.  This continues until there are no more piles left, at which point the
 * person with the most stones wins.
 *
 * Assuming Alex and Lee play optimally, return True if and only if Alex wins
 * the game.
 *
 *
 * Example 1:
 *
 *
 * Input: piles = [5,3,4,5]
 * Output: true
 * Explanation:
 * Alex starts first, and can only take the first 5 or the last 5.
 * Say he takes the first 5, so that the row becomes [3, 4, 5].
 * If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10
 * points.
 * If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win
 * with 9 points.
 * This demonstrated that taking the first 5 was a winning move for Alex, so we
 * return true.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= piles.length <= 500
 * piles.length is even.
 * 1 <= piles[i] <= 500
 * sum(piles) is odd.
 *
 *
 */
func stoneGame(piles []int) bool {
	// 1 2 3 4 5 7 8 5 4 3 2 1
	// 3 8 5 4 3 2 1 1
	// since the sum is odd, there'll always have a winner.
	// since number of piles is even, Alex always can arrange the remaining piles
	// to benefit him.
	return true
}
