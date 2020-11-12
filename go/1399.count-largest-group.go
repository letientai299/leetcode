package main

func countLargestGroup(n int) int {
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		s := 0
		t := i
		for t > 0 {
			s += t % 10
			t /= 10
		}
		m[s]++
	}

	s := 0
	max := 0
	for _, v := range m {
		if v > max {
			s = 1
			max = v
		} else if v == max {
			s++
		}
	}

	return s
}
