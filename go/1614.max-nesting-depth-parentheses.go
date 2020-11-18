package main

func maxDepth(s string) int {
	m := 0
	n := 0
	for _, c := range s {
		if c == '(' {
			n++
			if m < n {
				m = n
			}
		} else if c == ')' {
			n--
		}
	}
	return m
}
