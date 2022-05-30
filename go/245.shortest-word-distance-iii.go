package main

import "math"

// Shortest Word Distance III
//
// Medium
//
// https://leetcode.com/problems/shortest-word-distance-iii/
func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	x := -1
	best := math.MaxInt32
	for i, w := range wordsDict {
		if w == word1 || w == word2 {
			if x != -1 && (word1 == word2 || wordsDict[x] != w) {
				if i-x < best {
					best = i - x
				}
			}
			x = i
		}
	}

	return best
}
