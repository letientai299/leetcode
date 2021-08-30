package main

// Validate Stack Sequences
//
// Medium
//
// https://leetcode.com/problems/validate-stack-sequences/
//
// Given two integer arrays `pushed` and `popped` each with distinct values,
// return `true` _if this could have been the result of a sequence of push and
// pop operations on an initially empty stack, or_ `false` _otherwise._
//
// **Example 1:**
//
// ```
// Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// Output: true
// Explanation: We might do the following sequence:
// push(1), push(2), push(3), push(4),
// pop() -> 4,
// push(5),
// pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// Output: false
// Explanation: 1 cannot be popped before 2.
//
// ```
//
// **Constraints:**
//
// - `1 <= pushed.length <= 1000`
// - `0 <= pushed[i] <= 1000`
// - All the elements of `pushed` are **unique**.
// - `popped.length == pushed.length`
// - `popped` is a permutation of `pushed`.
func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stack := make([]int, 0, n)
	i, j := 0, 0
	for i < n && j < n {
		for i < n && (len(stack) == 0 || stack[len(stack)-1] != popped[j]) {
			stack = append(stack, pushed[i])
			i++
		}

		if len(stack) == 0 || stack[len(stack)-1] != popped[j] {
			return false
		}

		stack = stack[:len(stack)-1]
		j++
	}

	for k := len(stack) - 1; k >= 0 && j < n; k--{
		if stack[k]	 != popped[j] {
			return false
		}
		j++
	}

	return true
}

// Good for interview
