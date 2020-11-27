package main

// medium

func numberOfArithmeticSlices(a []int) int {
	n := len(a)
	if n < 3 {
		return 0
	}

	d := a[1] - a[0]
	cur := 2
	res := 0
	for i := 2; i < n; i++ {
		if a[i]-a[i-1] == d {
			cur++
			continue
		}

		for j := 1; cur-j >= 2; j++ {
			res += j
		}
		d = a[i] - a[i-1]
		cur = 2
	}

	for j := 1; cur-j >= 2; j++ {
		res += j
	}

	return res
}
