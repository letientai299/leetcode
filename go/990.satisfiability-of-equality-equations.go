package main

// Satisfiability of Equality Equations
//
// Medium
//
// https://leetcode.com/problems/satisfiability-of-equality-equations/
//
// You are given an array of strings `equations` that represent relationships
// between variables where each string `equations[i]` is of length `4` and takes
// one of two different forms: `"xi==yi"` or `"xi!=yi"`.Here, `xi` and `yi` are
// lowercase letters (not necessarily different) that represent one-letter
// variable names.
//
// Return `true` _if it is possible to assign integers to variable names so as
// to satisfy all the given equations, or_ `false` _otherwise_.
//
// **Example 1:**
//
// ```
// Input: equations = ["a==b","b!=a"]
// Output: false
// Explanation: If we assign say, a = 1 and b = 1, then the first equation is
// satisfied, but not the second.
// There is no way to assign the variables to satisfy both equations.
//
// ```
//
// **Example 2:**
//
// ```
// Input: equations = ["b==a","a==b"]
// Output: true
// Explanation: We could assign a = 1 and b = 1 to satisfy both equations.
//
// ```
//
// **Example 3:**
//
// ```
// Input: equations = ["a==b","b==c","a==c"]
// Output: true
//
// ```
//
// **Example 4:**
//
// ```
// Input: equations = ["a==b","b!=c","c==a"]
// Output: false
//
// ```
//
// **Example 5:**
//
// ```
// Input: equations = ["c==c","b==d","x!=z"]
// Output: true
//
// ```
//
// **Constraints:**
//
// - `1 <= equations.length <= 500`
// - `equations[i].length == 4`
// - `equations[i][0]` is a lowercase letter.
// - `equations[i][1]` is either `'='` or `'!'`.
// - `equations[i][2]` is `'='`.
// - `equations[i][3]` is a lowercase letter.
func equationsPossible(equations []string) bool {
	var groups [26]int
	numGroups := 1
	for _, f := range equations {
		x, y := f[0]-'a', f[3]-'a'
		eq := f[1] == '='

		if eq {
			if groups[x] == 0 && groups[y] == 0 {
				groups[x] = numGroups
				groups[y] = numGroups
				numGroups++
				continue
			}

			if groups[x] == 0 {
				groups[x] = groups[y]
				continue
			}

			if groups[y] == 0 {
				groups[y] = groups[x]
				continue
			}

			a, b := groups[x], groups[y]
			c := a
			if b < a {
				c = b
			}
			for i, v := range groups {
				if v == a || v == b {
					groups[i] = c
				}
			}
		}
	}

	for _, f := range equations {
		x, y := f[0]-'a', f[3]-'a'
		eq := f[1] == '='
		if !eq && ((x == y) || (groups[x] == groups[y] && groups[x] != 0)) {
			return false
		}
	}

	return true
}
