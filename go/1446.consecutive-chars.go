package main

func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := 0
	p := s[0]
	n := 1
	for i := 1; i < len(s); i++ {
		if p == s[i] {
			n++
			continue
		}

		if n > m {
			m = n
		}
		p = s[i]
		n = 1
	}
	if n > m {
		m = n
	}
	return m
}
