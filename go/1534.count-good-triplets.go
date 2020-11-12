package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	abs := func(x, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}

	n := len(arr)
	s := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i], arr[j]) > a {
				continue
			}

			for k := j + 1; k < n; k++ {
				if abs(arr[j], arr[k]) > b {
					continue
				}

				if abs(arr[i], arr[k]) > c {
					continue
				}

				s++
			}
		}
	}
	return s
}
