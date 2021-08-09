package main

import "math/rand"

// Implement Rand10() Using Rand7()
//
// Medium
//
// https://leetcode.com/problems/implement-rand10-using-rand7/
//
// Given the **API** `rand7()` that generates a uniform random integer in the
// range `[1, 7]`, write a function `rand10()` that generates a uniform random
// integer in the range `[1, 10]`. You can only call the API `rand7()`, and you
// shouldn't call any other API. Please **do not** use a language's built-in
// random API.
//
// Each test case will have one **internal** argument `n`, the number of times
// that your implemented function `rand10()` will be called while testing. Note
// that this is **not an argument** passed to `rand10()`.
//
// **Follow up:**
//
// - What is the [expected
// value](https://en.wikipedia.org/wiki/Expected_value) for the number of calls
// to `rand7()` function?
// - Could you minimize the number of calls to `rand7()`?
//
// **Example 1:**
//
// ```
// Input: n = 1
// Output: [2]
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 2
// Output: [2,8]
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 3
// Output: [3,8,10]
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 105`
func rand10() int {
	// idea from other people. I still don't understand Rejection Sampling
	r1 := rand7()
	for r1 == 7 {
		r1 = rand7()
	}

	r2 := rand7()
	for r2 > 5 {
		r2 = rand7()
	}

	return r2*2 - r1%2
}

func rand7() int {
	return rand.Intn(7) + 1
}
