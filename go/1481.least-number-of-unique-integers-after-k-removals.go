package main

import "sort"

// Least Number of Unique Integers after K Removals
//
// Medium
//
// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
//
// Given an array of integers `arr` and an integer `k`. Find the _least number
// of unique integers_ after removing **exactly** `k` elements **.**
//
// **Example 1:**
//
// ```
// Input: arr = [5,5,4], k = 1
// Output: 1
// Explanation: Remove the single 4, only 5 is left.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [4,3,1,1,3,3,2], k = 3
// Output: 2
// Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3
// will be left.
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 10^5`
// - `1 <= arr[i] <= 10^9`
// - `0 <= k <= arr.length`
func findLeastNumOfUniqueInts(arr []int, k int) int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	n := len(freq)
	freqCount := make(map[int]int)

	i := 0
	for _, f := range freq {
		if freqCount[f] == 0 {
			arr[i] = f
			i++
		}
		freqCount[f]++
	}

	arr = arr[:i]
	sort.Ints(arr)

	for i = 0; k > 0 && i < len(arr); i++ {
		f := arr[i]
		rem := freqCount[f] * f
		if rem > k {
			rem = k
			n -= k / f
		} else {
			n -= freqCount[f]
		}
		k -= rem
	}

	return n
}
