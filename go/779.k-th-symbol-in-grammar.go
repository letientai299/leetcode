package main

// K-th Symbol in Grammar
//
// Medium
//
// https://leetcode.com/problems/k-th-symbol-in-grammar/
//
// We build a table of `n` rows ( **1-indexed**). We start by writing `0` in the
// `1st` row. Now in every subsequent row, we look at the previous row and
// replace each occurrence of `0` with `01`, and each occurrence of `1` with
// `10`.
//
// - For example, for `n = 3`, the `1st` row is `0`, the `2nd` row is `01`, and
// the `3rd` row is `0110`.
//
// Given two integer `n` and `k`, return the `kth` ( **1-indexed**) symbol in
// the `nth` row of a table of `n` rows.
//
// **Example 1:**
//
// ```
// Input: n = 1, k = 1
// Output: 0
// Explanation: row 1: 0
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 2, k = 1
// Output: 0
// Explanation:
// row 1: 0
// row 2: 01
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 2, k = 2
// Output: 1
// Explanation:
// row 1: 0
// row 2: 01
//
// ```
//
// **Example 4:**
//
// ```
// Input: n = 3, k = 1
// Output: 0
// Explanation:
// row 1: 0
// row 2: 01
// row 3: 0110
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 30`
// - `1 <= k <= 2n - 1`
func kthGrammar(n int, k int) int {
	// 1: 0
	// 2: 01
	// 3: 01 10
	// 4: 0110 1001
	// 5: 01101001 10010110
	//
	// row n has 2^n bit, first bit is 0, last bit is (n-1)%2
	// that's why 1 <= k <= 2^(n-1)
	// bit(2x)   == 1-bit(x)
	// bit(2x-1) == 1-bit(2x) =  bit(x)

	x := 0
	for k > 1 {
		if k%2 == 0 {
			k /= 2
		} else {
			k++
		}
		x = 1 - x
	}

	return x
}
