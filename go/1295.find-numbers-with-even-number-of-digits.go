package main

func findNumbers(nums []int) int {
	n := 0
	for _, x := range nums {
		if x < 0 {
			x = -x
		}

		if (x >= 10 && x%100 == x) ||
			(x >= 1000 && x%10000 == x) ||
			(x >= 100000 && x%1000000 == x) ||
			(x >= 10000000 && x%100000000 == x) ||
			(x >= 1000000000 && x%10000000000 == x) {
			n++
		}
	}
	return n
}
