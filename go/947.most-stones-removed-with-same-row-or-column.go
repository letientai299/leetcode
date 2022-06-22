package main

// Most Stones Removed with Same Row or Column
//
// Medium
//
// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
//
// On a 2D plane, we place `n` stones at some integer coordinate points. Each
// coordinate point may have at most one stone.
//
// A stone can be removed if it shares either **the same row or the same
// column** as another stone that has not been removed.
//
// Given an array `stones` of length `n` where `stones[i] = [xi, yi]` represents
// the location of the `ith` stone, return _the largest possible number of
// stones that can be removed_.
//
// **Example 1:**
//
// ```
// Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
// Output: 5
// Explanation: One way to remove 5 stones is as follows:
// 1. Remove stone [2,2] because it shares the same row as [2,1].
// 2. Remove stone [2,1] because it shares the same column as [0,1].
// 3. Remove stone [1,2] because it shares the same row as [1,0].
// 4. Remove stone [1,0] because it shares the same column as [0,0].
// 5. Remove stone [0,1] because it shares the same row as [0,0].
// Stone [0,0] cannot be removed since it does not share a row/column with
// another stone still on the plane.
//
// ```
//
// **Example 2:**
//
// ```
// Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
// Output: 3
// Explanation: One way to make 3 moves is as follows:
// 1. Remove stone [2,2] because it shares the same row as [2,0].
// 2. Remove stone [2,0] because it shares the same column as [0,0].
// 3. Remove stone [0,2] because it shares the same row as [0,0].
// Stones [0,0] and [1,1] cannot be removed since they do not share a row/column
// with another stone still on the plane.
//
// ```
//
// **Example 3:**
//
// ```
// Input: stones = [[0,0]]
// Output: 0
// Explanation: [0,0] is the only stone on the plane, so you cannot remove it.
//
// ```
//
// **Constraints:**
//
// - `1 <= stones.length <= 1000`
// - `0 <= xi, yi <= 104`
// - No two stones are at the same coordinate point.
func removeStones(stones [][]int) int {
	// - 2 stones are connected if they share same row or columns
	// - In a connected path, we can remove all but one stone
	// - The problem reduces to counting connected paths

	rows := make(map[int]*int)
	cols := make(map[int]*int)
	pointers := make(map[int][]*int)
	for _, stone := range stones {
		r := stone[0]
		rp, rok := rows[r]
		c := stone[1]
		cp, cok := cols[c]
		if !rok && !cok {
			// new stone isn't belong to any path,
			// create a new path ID for it
			p := max(len(rows), len(cols)) + 1
			rows[r] = &p
			cols[c] = &p
			pointers[p] = append(pointers[p], &p)
			continue
		}

		if !rok {
			// new stone belong to known path identified via its column
			rows[r] = cp
			continue
		}

		if !cok {
			cols[c] = rp
			continue
		}

		// merge rp and cp, need to replace whole cp with rp, not just value
		// *cp = *rp // this won't work
		if *cp == *rp {
			continue
		}

		toBeDeletedID := *cp
		for _, p := range pointers[toBeDeletedID] {
			*p = *rp
			pointers[*rp] = append(pointers[*rp], p)
		}
		delete(pointers, toBeDeletedID)
	}

	uniq := make(map[int]bool)
	for _, p := range rows {
		uniq[*p] = true
	}
	for _, p := range cols {
		uniq[*p] = true
	}

	return len(stones) - len(uniq)
}

// TODO (tai): can be faster, and should find a better solution that don't
//  use many maps and pointers
