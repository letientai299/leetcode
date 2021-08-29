package main

// Beautiful Array
//
// Medium
//
// https://leetcode.com/problems/beautiful-array/
//
// An array `nums` of length `n` is **beautiful** if:
//
// - `nums` is a permutation of the integers in the range `[1, n]`.
// - For every `0 <= i < j < n`, there is no index `k` with `i < k < j` where `2
// * nums[k] == nums[i] + nums[j]`.
//
// Given the integer `n`, return _any **beautiful** array_ `nums` _of length_
// `n`. There will be at least one valid answer for the given `n`.
//
// **Example 1:**
//
// ```
// Input: n = 4
// Output: [2,1,4,3]
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 5
// Output: [3,1,2,5,4]
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 1000`
func beautifulArray(n int) []int {
	// Observations:
	//
	// - If we put all the even numbers in the first half, and odd numbers in the
	// second half, then for any 0<=i<n/2 and n/2<=j<n satisfy. We need to solve
	// the reduced problem for each half.
	// - Now, if we see those even numbers at 2x and odd numbers is 2x-1
	// where 0<x<n/2, we can use above logic to arrange them.

	res := make([]int, n)
	// fill inserts into the res array between l and r (exclusively)
	// those f(x) where x between 1 and end.
	var fill func(
		l, r int,
		f func(int) int, end int,
	)

	fill = func(l, r int, f func(int) int, end int) {
		if r-l <= 2 {
			for x := 1; x <= end; x++ {
				res[l] = f(x)
				l++
			}
			return
		}

		half := end / 2
		mid := l + half
		fill(l, mid, func(x int) int { return f(2 * x) }, half)
		fill(mid, r, func(x int) int { return f(2*x - 1) }, end-half)
	}

	fill(0, n, func(x int) int { return x }, n)
	return res
}
