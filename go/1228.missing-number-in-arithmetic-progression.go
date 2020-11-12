package main

func missingNumber(arr []int) int {
	n := len(arr)
	d1, d2 := arr[1]-arr[0], arr[n-1]-arr[n-2]
	if d1 > 0 {
		if d1 < d2 {
			return arr[n-1] - d1
		}

		if d1 > d2 {
			return arr[1] - d2
		}
	}

	if d1 < 0 {
		if d1 < d2 {
			return arr[0] + d2
		}

		if d1 > d2 {
			return arr[n-1] - d1
		}
	}

	for i := 2; i < n-1; i++ {
		if arr[i-1]+d1 != arr[i] {
			return arr[i-1] + d1
		}
	}

	return arr[0]
}
