package main

// Construct the Lexicographically Largest Valid Sequence
//
// Medium
//
// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
//
// Given an integer `n`, find a sequence that satisfies all of the following:
//
// - The integer `1` occurs once in the sequence.
// - Each integer between `2` and `n` occurs twice in the sequence.
// - For every integer `i` between `2` and `n`, the **distance** between the two
// occurrences of `i` is exactly `i`.
//
// The **distance** between two numbers on the sequence, `a[i]` and `a[j]`, is
// the absolute difference of their indices, `|j - i|`.
//
// Return _the **lexicographically largest** sequence_ _. It is guaranteed that
// under the given constraints, there is always a solution._
//
// A sequence `a` is lexicographically larger than a sequence `b` (of the same
// length) if in the first position where `a` and `b` differ, sequence `a` has a
// number greater than the corresponding number in `b`. For example, `[0,1,9,0]`
// is lexicographically larger than `[0,1,5,6]` because the first position they
// differ is at the third number, and `9` is greater than `5`.
//
// **Example 1:**
//
// ```
// Input: n = 3
// Output: [3,1,2,3,2]
// Explanation: [2,3,2,1,3] is also a valid sequence, but [3,1,2,3,2] is the
// lexicographically largest valid sequence.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 5
// Output: [5,3,1,4,3,5,2,4,2]
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 20`
func constructDistancedSequence(n int) []int {
	res := make([]int, 2*n-1)
	used := 0
	var walk func(i int) bool

	walk = func(i int) bool {
		for i < len(res) && res[i] != 0 {
			i++
		}

		if i >= len(res) {
			return used != (1<<(n+1)-1)
		}

		v := n
		for v > 0 {
			next := i + v
			if v == 1 {
				next--
			}

			if used&(1<<v) != 0 || next >= len(res) || res[next] != 0 {
				v--
				continue
			}

			res[i] = v
			res[next] = v
			used |= 1 << v
			if walk(i + 1) {
				return true
			}
			used &= ^(1 << v)
			res[i] = 0
			res[next] = 0
			v--
		}

		return false
	}

	walk(0)
	return res
}
