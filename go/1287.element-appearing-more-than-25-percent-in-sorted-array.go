package main

func findSpecialInteger(arr []int) int {
	if len(arr) < 4 {
		return arr[0]
	}

	n := len(arr)
	prev := arr[0]
	count := 1
	for i := 1; i < n; i++ {
		if arr[i] != prev {
			prev = arr[i]
			count = 1
			continue
		}

		count++
		if count > n/4 {
			return prev
		}
	}

	return prev
}
