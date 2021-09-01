package main

// Max Chunks To Make Sorted
//
// Medium
//
// https://leetcode.com/problems/max-chunks-to-make-sorted/
//
// You are given an integer array `arr` of length `n` that represents a
// permutation of the integers in the range `[0, n - 1]`.
//
// We split `arr` into some number of **chunks** (i.e., partitions), and
// individually sort each chunk. After concatenating them, the result should
// equal the sorted array.
//
// Return _the largest number of chunks we can make to sort the array_.
//
// **Example 1:**
//
// ```
// Input: arr = [4,3,2,1,0]
// Output: 1
// Explanation:
// Splitting into two or more chunks will not return the required result.
// For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2],
// which isn't sorted.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [1,0,2,3,4]
// Output: 4
// Explanation:
// We can split into two chunks, such as [1, 0], [2, 3, 4].
// However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks
// possible.
//
// ```
//
// **Constraints:**
//
// - `n == arr.length`
// - `1 <= n <= 10`
// - `0 <= arr[i] < n`
// - All the elements of `arr` are **unique**.
func maxChunksToSorted(arr []int) int {
	var ranges [][]int
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if i <= v {
			ranges = append(ranges, []int{i, v})
			continue
		}

		// new range is [v, i]
		j := 0
		for _, pre := range ranges {
			if pre[1] < v {
				ranges[j] = pre
				j++
				continue // disjoint
			}

			if pre[0] > v {
				pre[0] = v
			} else {
				v = pre[0]
			}

			pre[1] = i
		}
		ranges[j][0] = v
		ranges[j][1] = i
		ranges = ranges[:j+1]
	}

	return len(ranges)
}

// TODO (tai): this is O(n^2) and O(n) space, try again with O(n) time and O(1) space
