package main

// Finding Pairs With a Certain Sum
//
// Medium
//
// https://leetcode.com/problems/finding-pairs-with-a-certain-sum/
//
// You are given two integer arrays `nums1` and `nums2`. You are tasked to
// implement a data structure that supports queries of two types:
//
// 1. **Add** a positive integer to an element of a given index in the array
// `nums2`.
// 2. **Count** the number of pairs `(i, j)` such that `nums1[i] + nums2[j]`
// equals a given value (`0 <= i < nums1.length` and `0 <= j < nums2.length`).
//
// Implement the `FindSumPairs` class:
//
// - `FindSumPairs(int[] nums1, int[] nums2)` Initializes the `FindSumPairs`
// object with two integer arrays `nums1` and `nums2`.
// - `void add(int index, int val)` Adds `val` to `nums2[index]`, i.e., apply
// `nums2[index] += val`.
// - `int count(int tot)` Returns the number of pairs `(i, j)` such that
// `nums1[i] + nums2[j] == tot`.
//
// **Example 1:**
//
// ```
// ,5,42,5,42,5,4
// ```
//
// **Constraints:**
//
// - `1 <= nums1.length <= 1000`
// - `1 <= nums2.length <= 105`
// - `1 <= nums1[i] <= 109`
// - `1 <= nums2[i] <= 105`
// - `0 <= index < nums2.length`
// - `1 <= val <= 105`
// - `1 <= tot <= 109`
// - At most `1000` calls are made to `add` and `count` **each**.

type FindSumPairs struct {
	m     map[int]int
	nums2 []int
	nums1 map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	sp := FindSumPairs{
		m:     make(map[int]int, len(nums2)),
		nums1: make(map[int]int, len(nums1)),
		nums2: nums2,
	}

	for _, y := range nums1 {
		sp.nums1[y]++
	}

	for _, x := range nums2 {
		sp.m[x]++
	}

	return sp
}

func (sp *FindSumPairs) Add(index int, val int) {
	old := sp.nums2[index]
	sp.m[old]--
	cur := old + val
	sp.m[cur]++
	sp.nums2[index] = cur
}

func (sp *FindSumPairs) Count(tot int) int {
	r := 0
	for x, cx := range sp.nums1 {
		y := tot - x
		r += sp.m[y] * cx
	}
	return r
}
