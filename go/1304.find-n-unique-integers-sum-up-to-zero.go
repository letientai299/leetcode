package main

func sumZero(n int) []int {
	a := make([]int, 0, n)
	if n%2 == 1 {
		a = append(a, 0)
		n--
	}

	for n > 0 {
		a = append(a, n, -n)
		n -= 2
	}

	return a
}
