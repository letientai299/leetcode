package main

import "sort"

// Reduce Array Size to The Half
//
// Medium
//
// https://leetcode.com/problems/reduce-array-size-to-the-half/
//
// You are given an integer array `arr`. You can choose a set of integers and
// remove all the occurrences of these integers in the array.
//
// Return _the minimum size of the set so that **at least** half of the integers
// of the array are removed_.
//
// **Example 1:**
//
// ```
// Input: arr = [3,3,3,3,5,5,5,2,2,7]
// Output: 2
// Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has
// size 5 (i.e equal to half of the size of the old array).
// Possible sets of size 2 are {3,5},{3,2},{5,2}.
// Choosing set {2,7} is not possible as it will make the new array
// [3,3,3,3,5,5,5] which has size greater than half of the size of the old
// array.
//
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [7,7,7,7,7,7]
// Output: 1
// Explanation: The only possible set you can choose is {7}. This will make the
// new array empty.
//
// ```
//
// **Example 3:**
//
// ```
// Input: arr = [1,9]
// Output: 1
//
// ```
//
// **Example 4:**
//
// ```
// Input: arr = [1000,1000,3,7]
// Output: 1
//
// ```
//
// **Example 5:**
//
// ```
// Input: arr = [1,2,3,4,5,6,7,8,9,10]
// Output: 5
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 105`
// - `arr.length` is even.
// - `1 <= arr[i] <= 105`
func minSetSize(arr []int) int {
	n := len(arr)
	m := make(map[int]int, len(arr)/2)
	for _, v := range arr {
		m[v]++
	}
	i := 0
	for _, v := range m {
		arr[i] = v
		i++
	}

	arr = arr[:i]
	sort.Ints(arr)
	sum := 0
	j := i - 1
	for ; j >= 0 && sum < (n+1)/2; j-- {
		sum += arr[j]
	}

	return i - j - 1
}
