package main

// Decode XORed Permutation
//
// Medium
//
// https://leetcode.com/problems/decode-xored-permutation/
//
// There is an integer array `perm` that is a permutation of the first `n`
// positive integers, where `n` is always **odd**.
//
// It was encoded into another integer array `encoded` of length `n - 1`, such
// that `encoded[i] = perm[i] XOR perm[i + 1]`. For example, if `perm =
// [1,3,2]`, then `encoded = [2,1]`.
//
// Given the `encoded` array, return _the original array_ `perm`. It is
// guaranteed that the answer exists and is unique.
//
// **Example 1:**
//
// ```
// Input: encoded = [3,1]
// Output: [1,2,3]
// Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
//
// ```
//
// **Example 2:**
//
// ```
// Input: encoded = [6,5,4,6]
// Output: [2,4,1,5,3]
//
// ```
//
// **Constraints:**
//
// - `3 <= n < 105`
// - `n` is odd.
// - `encoded.length == n - 1`
func decode(encoded []int) []int {
	n := len(encoded) + 1
	all := 0
	for i := 1; i <= n; i++ {
		all ^= i
	}

	e := 0
	for i := 1; i < len(encoded); i += 2 {
		e ^= encoded[i]
	}

	first := all ^ e

	res := make([]int, n)
	res[0] = first

	for i := 1; i < n; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}
