package main

// Letter Tile Possibilities
//
// Medium
//
// https://leetcode.com/problems/letter-tile-possibilities/
//
// You have `n``tiles`, where each tile has one letter `tiles[i]` printed on it.
//
// Return _the number of possible non-empty sequences of letters_ you can make
// using the letters printed on those `tiles`.
//
// **Example 1:**
//
// ```
// Input: tiles = "AAB"
// Output: 8
// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB",
// "ABA", "BAA".
//
// ```
//
// **Example 2:**
//
// ```
// Input: tiles = "AAABBC"
// Output: 188
//
// ```
//
// **Example 3:**
//
// ```
// Input: tiles = "V"
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= tiles.length <= 7`
// - `tiles` consists of uppercase English letters.
func numTilePossibilities(tiles string) int {
	var m [26]int
	for _, c := range tiles {
		m[c-'A']++
	}
	var uniq []int
	for _, v := range m {
		if v > 0 {
			uniq = append(uniq, v)
		}
	}

	r := 0

	fac := func(x int) int {
		v := 1
		for x > 0 {
			v *= x
			x--
		}
		return v
	}

	n := len(uniq)
	uses := make([]int, n)
	calc := func() int {
		total := 0
		p := 1
		for _, v := range uses {
			total += v
			p *= fac(v)
		}
		if total == 0 {
			return 0
		}

		return fac(total) / p
	}

	var walk func(i int)
	walk = func(i int) {
		if i >= len(uniq) {
			r += calc()
			return
		}

		for j := 0; j <= uniq[i]; j++ {
			uses[i] = j
			walk(i + 1)
		}
	}

	walk(0)
	return r
}
