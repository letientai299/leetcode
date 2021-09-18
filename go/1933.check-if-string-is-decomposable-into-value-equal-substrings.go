package main

// Check if String Is Decomposable Into Value-Equal Substrings
//
// Easy
//
// https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/
func isDecomposable(s string) bool {
	n := len(s)
	sub2 := false
	i := 0
	for ; i < n-2; i += 3 {
		if s[i] != s[i+1] {
			return false
		}

		if s[i+1] == s[i+2] {
			continue
		}

		if sub2 {
			return false
		}

		sub2 = true
		i--
	}

	if i == n-2 && !sub2 {
		sub2 = s[i] == s[i+1]
		i += 2
	}

	return sub2 && i == n // must exists sub string of len 2
}
