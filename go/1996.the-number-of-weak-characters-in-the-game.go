package main

import (
	"sort"
)

// The Number of Weak Characters in the Game
//
// Medium
//
// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
//
// You are playing a game that contains multiple characters, and each of the
// characters has **two** main properties: **attack** and **defense**. You are
// given a 2D integer array `properties` where `properties[i] = [attacki,
// defensei]` represents the properties of the `ith` character in the game.
//
// A character is said to be **weak** if any other character has **both** attack
// and defense levels **strictly greater** than this character's attack and
// defense levels. More formally, a character `i` is said to be **weak** if
// there exists another character `j` where `attackj > attacki` and `defensej >
// defensei`.
//
// Return _the number of **weak** characters_.
//
// **Example 1:**
//
// ```
// Input: properties = [[5,5],[6,3],[3,6]]
// Output: 0
// Explanation: No character has strictly greater attack and defense than the
// other.
//
// ```
//
// **Example 2:**
//
// ```
// Input: properties = [[2,2],[3,3]]
// Output: 1
// Explanation: The first character is weak because the second character has a
// strictly greater attack and defense.
//
// ```
//
// **Example 3:**
//
// ```
// Input: properties = [[1,5],[10,4],[4,3]]
// Output: 1
// Explanation: The third character is weak because the second character has a
// strictly greater attack and defense.
//
// ```
//
// **Constraints:**
//
// - `2 <= properties.length <= 105`
// - `properties[i].length == 2`
// - `1 <= attacki, defensei <= 105`
func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	atk := make([]int, n)
	for i := range atk {
		atk[i] = i
	}

	sort.Slice(atk, func(i, j int) bool {
		a, b := properties[atk[i]], properties[atk[j]]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] > b[0]
	})

	r := 0
	cur := properties[atk[0]]
	for i := 1; i < n; i++ {
		p := properties[atk[i]]
		if p[1] < cur[1] {
			if p[0] < cur[0] {
				r++
			}
		} else {
			if p[1] > cur[1] {
				cur = p
			}
		}
	}
	return r
}
