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
