package main

// Find Kth Bit in Nth Binary String
//
// Medium
//
// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/
//
// Given two positive integers `n` and `k`, the binary string  `Sn` is formed as
// follows:
//
// - `S1 = "0"`
// - `Si = Si-1 + "1" + reverse(invert(Si-1))` for `i > 1`
//
// Where `+` denotes the concatenation operation, `reverse(x)` returns the
// reversed string x, and `invert(x)` inverts all the bits in x (0 changes to 1
// and 1 changes to 0).
//
// For example, the first 4 strings in the above sequence are:
//
// - `S1 = "0"`
// - `S2 = "011"`
// - `S3 = "0111001"`
// - `S4 = "011100110110001"`
//
// Return _the_ `kth` _bit_ _in_ `Sn`. It is guaranteed that `k` is valid for
// the given `n`.
//
// **Example 1:**
//
// ```
// Input: n = 3, k = 1
// Output: "0"
// Explanation: S3 is "0111001". The first bit is "0".
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 4, k = 11
// Output: "1"
// Explanation: S4 is "011100110110001". The 11th bit is "1".
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 1, k = 1
// Output: "0"
//
// ```
//
// **Example 4:**
//
// ```
// Input: n = 2, k = 3
// Output: "1"
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 20`
// - `1 <= k <= 2n - 1`
func findKthBit(n int, k int) byte {
	if n == 1 || k == 1 {
		return '0'
	}

	// len(S(n)) == 2^n-1
	// S(n)[mid] == 1
	// S(n)[0] == 0
	// S(n)[-1] == 1 or 0 (if n == 1)
	mid := 1 << (n - 1) // k is 1-indexed

	if k == mid {
		return '1'
	}

	if k < mid {
		return findKthBit(n-1, k)
	}

	return 1 - (findKthBit(n-1, 2*mid-k) - '0') + '0'
}

// Good for interview
