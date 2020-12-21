package main

// medium

func countTriplets(arr []int) int {
	n := len(arr)
	var cnt int
	for i, v := range arr {
		for j := i + 1; j < n; j++ {
			v ^= arr[j]
			if v == 0 {
				cnt += j - i
			}
		}
	}

	return cnt
}
