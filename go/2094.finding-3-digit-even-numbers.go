package main

import "sort"

// Finding 3-Digit Even Numbers
//
// Easy
//
// https://leetcode.com/problems/finding-3-digit-even-numbers/
//
// You are given an integer array `digits`, where each element is a digit. The
// array may contain duplicates.
//
// You need to find **all** the **unique** integers that follow the given
// requirements:
//
// - The integer consists of the **concatenation** of **three** elements from
// `digits` in **any** arbitrary order.
// - The integer does not have **leading zeros**.
// - The integer is **even**.
//
// For example, if the given `digits` were `[1, 2, 3]`, integers `132` and `312`
// follow the requirements.
//
// Return _a **sorted** array of the unique integers._
//
// **Example 1:**
//
// ```
// Input: digits = [2,1,3,0]
// Output: [102,120,130,132,210,230,302,310,312,320]
// Explanation: All the possible integers that follow the requirements are in
// the output array.
// Notice that there are no odd integers or integers with leading zeros.
//
// ```
//
// **Example 2:**
//
// ```
// Input: digits = [2,2,8,8,2]
// Output: [222,228,282,288,822,828,882]
// Explanation: The same digit can be used as many times as it appears in
// digits.
// In this example, the digit 8 is used twice each time in 288, 828, and 882.
//
// ```
//
// **Example 3:**
//
// ```
// Input: digits = [3,7,5]
// Output: []
// Explanation: No even integers can be formed using the given digits.
//
// ```
//
// **Constraints:**
//
// - `3 <= digits.length <= 100`
// - `0 <= digits[i] <= 9`
func findEvenNumbers(digits []int) []int {
	if len(digits) < 3 {
		return nil
	}

	sort.Ints(digits)

	even := make([]int, 0, len(digits)/2)
	for _, d := range digits {
		if d%2 == 0 {
			even = append(even, d)
		}
	}

	if len(even) == 0 {
		return nil // no even digits
	}

	var r []int

	used := make(map[int]bool, len(digits))
	seen := make(map[int]bool)
	var gen func(cur, x10 int)

	gen = func(cur, x10 int) {
		if x10 > 2 {
			if !seen[cur] {
				r = append(r, cur)
				seen[cur] = true
			}
			return
		}

		for i, d := range digits {
			if used[i] {
				continue
			}

			if x10 == 2 && d%2 == 1 {
				continue
			}

			if x10 == 0 && d == 0 {
				continue
			}

			used[i] = true
			gen(cur*10+d, x10+1)
			used[i] = false
		}
	}

	gen(0, 0)
	return r

	// TODO (tai): can be faster, there's O(n) or O(900) solution
}
