package main

// Possible Bipartition
//
// Medium
//
// https://leetcode.com/problems/possible-bipartition/
//
// We want to split a group of `n` people (labeled from `1` to `n`) into two
// groups of **any size**. Each person may dislike some other people, and they
// should not go into the same group.
//
// Given the integer `n` and the array `dislikes` where `dislikes[i] = [ai, bi]`
// indicates that the person labeled `ai` does not like the person labeled `bi`,
// return `true` _if it is possible to split everyone into two groups in this
// way_.
//
// **Example 1:**
//
// ```
// Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
// Output: true
// Explanation: group1 [1,4] and group2 [2,3].
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 3, dislikes = [[1,2],[1,3],[2,3]]
// Output: false
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 2000`
// - `0 <= dislikes.length <= 104`
// - `dislikes[i].length == 2`
// - `1 <= dislikes[i][j] <= n`
// - `ai < bi`
// - All the pairs of `dislikes` are **unique**.
func possibleBipartition(n int, dislikes [][]int) bool {
	// init groups where all is 0, means doesn't check yet
	groups := make([]int8, n)

	hate := make(map[int][]int)
	for _, d := range dislikes {
		a, b := d[0]-1, d[1]-1
		hate[a] = append(hate[a], b)
		hate[b] = append(hate[b], a)
	}

	var walk func(i int, sign int8) bool

	walk = func(i int, me int8) bool {
		for _, x := range hate[i] {
			other := groups[x]
			if other == 0 {
				continue
			}

			if me == 0 && other != 0 {
				me = -other
				break
			}

			if me*other != -1 {
				return false
			}
		}

		if me == 0 {
			me = 1
		}

		groups[i] = me
		for _, v := range hate[i] {
			o := groups[v]
			if o*me > 0 {
				return false
			}

			if o != 0 {
				continue
			}

			groups[v] = -me
			if !walk(v, -me) {
				return false
			}
		}
		return true
	}

	for i := range hate {
		if groups[i] == 0 {
			if !walk(i, 0) {
				return false
			}
		}
	}

	return true
}

// TODO (tai): can be faster
