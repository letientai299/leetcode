package main

// Minimum Number of Frogs Croaking
//
// Medium
//
// https://leetcode.com/problems/minimum-number-of-frogs-croaking/
//
// Given the string `croakOfFrogs`, which represents a combination of the string
// "croak" from different frogs, that is, multiple frogs can croak at the same
// time, so multiple “croak” are mixed. _Return the minimum number of_ different
// _frogs to finish all the croak in the given string._
//
// A valid "croak" means a frog is printing 5 letters ‘c’, ’r’, ’o’, ’a’,
// ’k’ **sequentially**. The frogs have to print all five letters to finish a
// croak. If the given string is not a combination of valid "croak" return -1.
//
// **Example 1:**
//
// ```
// Input: croakOfFrogs = "croakcroak"
// Output: 1
// Explanation: One frog yelling "croak" twice.
//
// ```
//
// **Example 2:**
//
// ```
// Input: croakOfFrogs = "crcoakroak"
// Output: 2
// Explanation: The minimum number of frogs is two.
// The first frog could yell "crcoakroak".
// The second frog could yell later "crcoakroak".
//
// ```
//
// **Example 3:**
//
// ```
// Input: croakOfFrogs = "croakcrook"
// Output: -1
// Explanation: The given string is an invalid combination of "croak" from
// different frogs.
//
// ```
//
// **Example 4:**
//
// ```
// Input: croakOfFrogs = "croakcroa"
// Output: -1
//
// ```
//
// **Constraints:**
//
// - `1 <= croakOfFrogs.length <= 10^5`
// - All characters in the string are: `'c'`, `'r'`, `'o'`, `'a'` or `'k'`.
func minNumberOfFrogs(s string) int {
	m := map[int32]int{
		'c': 0,
		'r': 1,
		'o': 2,
		'a': 3,
		'k': 4,
	}

	best := 0
	var counts [5]int
	for _, c := range s {
		i := m[c]
		if i > 0 && counts[i] >= counts[i-1] {
			return -1
		}
		counts[i]++

		if c == 'k' {
			max := 0
			for j, v := range counts {
				if v > max {
					max = v
				}
				counts[j] = v - 1
			}

			if best < max {
				best = max
			}
		}
	}

	for _, v := range counts {
		if v != 0 {
			return -1
		}
	}

	return best
}
