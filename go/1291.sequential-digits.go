package main

// Sequential Digits
//
// Medium
//
// https://leetcode.com/problems/sequential-digits/
//
// An integer has _sequential digits_ if and only if each digit in the number is
// one more than the previous digit.
//
// Return a **sorted** list of all the integers in the range `[low,
// high]` inclusive that have sequential digits.
//
// **Example 1:**
//
// ```
// Input: low = 100, high = 300
// Output: [123,234]
//
// ```
//
// **Example 2:**
//
// ```
// Input: low = 1000, high = 13000
// Output: [1234,2345,3456,4567,5678,6789,12345]
//
// ```
//
// **Constraints:**
//
// - `10 <= low <= high <= 10^9`
func sequentialDigits(low int, high int) []int {
	var res []int
	v := 0
	d := 1
	inc := 0
	ten := 1

	for t := low; t > 0; {
		v = 10*v + d
		d++
		t /= 10
		inc = 10*inc + 1
		ten *= 10
	}

	for {
		for x := v; x < ten; {
			if x >= low && x <= high {
				res = append(res, x)
			}
			if x%10 == 9 {
				break
			}
			x += inc
		}
		ten *= 10
		v = 10*v + d
		d++
		inc = 10*inc + 1
		if ten > high*10 {
			break
		}
	}

	return res
}
