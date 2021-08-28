package main

import (
	"strings"
)

// Permutation Sequence
//
// Hard
//
// https://leetcode.com/problems/permutation-sequence/
//
// The set `[1, 2, 3, ..., n]` contains a total of `n!` unique permutations.
//
// By listing and labeling all of the permutations in order, we get the
// following sequence for `n = 3`:
//
// 1. `"123"`
// 2. `"132"`
// 3. `"213"`
// 4. `"231"`
// 5. `"312"`
// 6. `"321"`
//
// Given `n` and `k`, return the `kth` permutation sequence.
//
// **Example 1:**
//
// ```
// Input: n = 3, k = 3
// Output: "213"
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 4, k = 9
// Output: "2314"
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 3, k = 1
// Output: "123"
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 9`
// - `1 <= k <= n!`
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	factor := 1
	for i := range nums {
		nums[i] = i + 1
		factor *= i + 1
	}

	var sb strings.Builder
	for x := n; x > 0; x-- {
		factor /= x
		i := (k - 1) / factor
		d := nums[i]
		nums = append(nums[:i], nums[i+1:]...)
		sb.WriteByte(byte(d + '0'))
		k = (k-1)%factor + 1
	}

	return sb.String()
}

func getPermutationBrute(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	for k > 1 {
		k--
		for i := n - 1; i >= 1; i-- {
			if nums[i] < nums[i-1] {
				continue
			}

			i--
			j := n - 1
			for ; j > i; j-- {
				if nums[j] > nums[i] {
					break
				}
			}

			nums[i], nums[j] = nums[j], nums[i]
			i++
			j = n - 1
			for i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
			break
		}
	}

	var sb strings.Builder
	for _, v := range nums {
		sb.WriteByte(byte(v + '0'))
	}

	return sb.String()
}
