package main

// medium

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 0
	a, b := s[0], s[0]
	start := 0
	for i := 1; i < len(s); i++ {
		x := s[i]
		if x == a || x == b {
			continue
		}

		if x != a && a == b {
			b = x
			continue
		}

		if x != a && x != b {
			if max < i-start {
				max = i - start
			}

			start = i - 1
			for start >= 0 && s[start] == s[i-1] {
				start--
			}
			start++
			a, b = s[i-1], s[i]
		}
	}

	if max < len(s)-start {
		max = len(s) - start
	}
	return max
}
