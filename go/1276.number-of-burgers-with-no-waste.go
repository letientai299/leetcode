package main

// medium

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// x + y = b
	// 2x + 2y = 2b
	// 4x + 2y = a

	n := tomatoSlices - 2*cheeseSlices
	if tomatoSlices%2 != 0 ||
		n < 0 ||
		n%2 != 0 ||
		n/2 > cheeseSlices {
		return nil
	}

	return []int{n / 2, cheeseSlices - n/2}
}
