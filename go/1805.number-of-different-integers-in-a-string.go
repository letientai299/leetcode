package main

// Number of Different Integers in a String
//
// Easy
//
// https://leetcode.com/problems/number-of-different-integers-in-a-string/
//
// You are given a string `word` that consists of digits and lowercase English
// letters.
//
// You will replace every non-digit character with a space. For example,
// `"a123bc34d8ef34"` will become `" 123  34 8  34"`. Notice that you are left
// with some integers that are separated by at least one space: `"123"`, `"34"`,
// `"8"`, and `"34"`.
//
// Return _the number of **different** integers after performing the replacement
// operations on_ `word`.
//
// Two integers are considered different if their decimal representations
// **without any leading zeros** are different.
//
// **Example 1:**
//
// ```
// Input: word = "a123bc34d8ef34"
// Output: 3
// Explanation: The three different integers are "123", "34", and "8". Notice
// that "34" is only counted once.
//
// ```
//
// **Example 2:**
//
// ```
// Input: word = "leet1234code234"
// Output: 2
//
// ```
//
// **Example 3:**
//
// ```
// Input: word = "a1b01c001"
// Output: 1
// Explanation: The three integers "1", "01", and "001" all represent the same
// integer because
// the leading zeros are ignored when comparing their decimal values.
//
// ```
//
// **Constraints:**
//
// - `1 <= word.length <= 1000`
// - `word` consists of digits and lowercase English letters.
func numDifferentIntegers(word string) int {
	// store unique numbers as string, without leading 0
	m := make(map[string]struct{})
	start := -1 // not running in a number yet

	trim := func(end int) int {
		for i := start; i < end; i++ {
			if word[i] != '0' {
				return i
			}
		}
		return end - 1
	}

	for i, c := range word {
		if c >= '0' && c <= '9' {
			if start == -1 {
				start = i // start of a number
			}
			continue
		}

		if start == -1 {
			continue
		}

		// number ends with previous char
		start = trim(i)
		m[word[start:i]] = struct{}{}
		start = -1
	}

	if start != -1 {
		start = trim(len(word))
		m[word[start:]] = struct{}{}
	}

	return len(m)
}

// Good for interview
