package main

import (
	"math"
	"sort"
)

// Sum of Mutated Array Closest to Target
//
// Medium
//
// https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/
//
// Given an integer array `arr` and a target value `target`, return the integer
// `value` such that when we change all the integers larger than `value` in the
// given array to be equal to `value`, the sum of the array gets as close as
// possible (in absolute difference) to `target`.
//
// In case of a tie, return the minimum such integer.
//
// Notice that the answer is not neccesarilly a number from `arr`.
//
// **Example 1:**
//
// ```
// Input: arr = [4,9,3], target = 10
// Output: 3
// Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's
// the optimal answer.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [2,3,5], target = 10
// Output: 5
//
// ```
//
// **Example 3:**
//
// ```
// Input: arr = [60864,25176,27249,21296,20204], target = 56803
// Output: 11361
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 104`
// - `1 <= arr[i], target <= 105`
func findBestValue(arr []int, target int) int {
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	sort.Ints(arr)
	s := 0
	n := len(arr)
	best := arr[n-1]
	diff := math.MaxInt32

	arr = append(arr, 0)
	for i, v := range arr {
		t := target - s
		count := n - i

		if count == 0 {
			s += v
			continue
		}

		x := (target - s) / count
		if i > 0 && x < arr[i-1] {
			break
		}

		if x > v {
			s += v
			continue
		}

		y := x + 1
		a := abs(t - x*count)
		b := abs(t - y*count)

		if b < a {
			a = b
			x = y
		}

		if a < diff {
			diff = a
			best = x
		} else if a == diff && best > x {
			best = x
		}

		s += v
	}

	return best
}
