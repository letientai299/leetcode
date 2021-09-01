package main

// Minimum Domino Rotations For Equal Row
//
// Medium
//
// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
//
// In a row of dominoes, `tops[i]` and `bottoms[i]` represent the top and bottom
// halves of the `ith` domino. (A domino is a tile with two numbers from 1 to 6
// - one on each half of the tile.)
//
// We may rotate the `ith` domino, so that `tops[i]` and `bottoms[i]` swap
// values.
//
// Return the minimum number of rotations so that all the values in `tops` are
// the same, or all the values in `bottoms` are the same.
//
// If it cannot be done, return `-1`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/05/14/domino.png)
//
// ```
// Input: tops = [2,1,2,4,2,2], bottoms = [5,2,6,2,3,2]
// Output: 2
// Explanation:
// The first figure represents the dominoes as given by tops and bottoms: before
// we do any rotations.
// If we rotate the second and fourth dominoes, we can make every value in the
// top row equal to 2, as indicated by the second figure.
//
// ```
//
// **Example 2:**
//
// ```
// Input: tops = [3,5,1,2,3], bottoms = [3,6,3,3,4]
// Output: -1
// Explanation:
// In this case, it is not possible to rotate the dominoes to make one row of
// values equal.
//
// ```
//
// **Constraints:**
//
// - `2 <= tops.length <= 2 * 104`
// - `bottoms.length == tops.length`
// - `1 <= tops[i], bottoms[i] <= 6`
func minDominoRotations(tops []int, bottoms []int) int {
	var countTops [6]int
	var countBots [6]int
	var both [6]int
	n := len(tops)
	for i, t := range tops {
		b := bottoms[i]
		if b == t {
			both[b-1]++
		} else {
			countTops[t-1]++
			countBots[b-1]++
		}
	}

	best := -1
	for i, t := range countTops {
		b := countBots[i]
		same := both[i]

		if same+b+t < n {
			continue
		}

		b = n - b - same
		if b > n-t-same {
			b = n - t - same
		}

		if best == -1 || best > b {
			best = b
		}
	}

	return best
}

// Good for interview
