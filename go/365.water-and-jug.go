package main

// medium

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}

	if x+y == 0 {
		return z == 0
	}

	c := gcd(x, y)
	return z%c == 0
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	if a == 1 {
		return 1
	}

	return gcd(b%a, a)
}
