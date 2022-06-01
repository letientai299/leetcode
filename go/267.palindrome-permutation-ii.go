package main

// Palindrome Permutation II
//
// Medium
//
// https://leetcode.com/problems/palindrome-permutation-ii/
func generatePalindromes(s string) []string {
	// count
	odds := 0
	var m = [26]int{}
	for _, c := range s {
		c -= 'a'
		m[c]++
		if m[c]%2 == 0 {
			odds--
		} else {
			odds++
		}
	}

	if odds > 1 {
		return nil
	}

	mid := (len(s) - odds) / 2
	bs := make([]byte, len(s))
	for i, v := range m {
		if v%2 == 1 {
			c := i + 'a'
			bs[mid] = byte(c)
		}

		m[i] = v / 2
	}

	// generate all half strings.
	var r []string

	var gen func(i int)
	gen = func(i int) {
		if i == mid {
			r = append(r, string(bs))
			return
		}

		for j, k := range m {
			if k > 0 {
				m[j]--
				c := byte(j + 'a')
				bs[i] = c
				bs[len(s)-1-i] = c
				gen(i + 1)
				m[j]++
			}
		}

	}

	gen(0)
	return r
}
