package main

// Sum Multiples
//
// Easy
//
// https://leetcode.com/problems/sum-multiples
//
// Given a positive integer `n`, find the sum of all integers in the range `[1,
// n]` **inclusive** that are divisible by `3`, `5`, or `7`.
//
// Return _an integer denoting the sum of all numbers in the given range
// satisfying the constraint._
//
// **Example 1:**
//
// ```
// [1, 7]35,7 3, 5, 6, 721
// ```
//
// **Example 2:**
//
// ```
// [1, 10] that are35,73, 5, 6, 7, 9, 10
// ```
//
// **Example 3:**
//
// ```
// [1, 9]3573, 5, 6, 7, 930
// ```
//
// **Constraints:**
//
// - `1 <= n <= 103`
func sumOfMultiples(n int) int {
  r := 0
  for i := range n {
    x := i + 1
    if x%3 == 0 || x%5 == 0 || x%7 == 0 {
      r += x
    }
  }
  return r
}
