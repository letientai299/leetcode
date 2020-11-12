package main

func getNoZeroIntegers(n int) []int {
	noZero := func(x int) bool {
		for x > 0 {
			if x%10 == 0 {
				return false
			}
			x /= 10
		}
		return true
	}

	for x := 1; x <= n/2; x++ {
		if !noZero(x) {
			continue
		}

		y := n - x
		if noZero(y) {
			return []int{x, y}
		}
	}

	return nil
}
