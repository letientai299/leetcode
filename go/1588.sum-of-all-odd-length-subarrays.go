package main

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return arr[0]
	}

	for i := 1; i < n; i++ {
		arr[i] += arr[i-1]
	}

	s := 0
	arr = append([]int{0}, arr...)
	for d := 1; d <= n; d += 2 {
		for i := 0; i <= n-d; i++ {
			s += arr[i+d] - arr[i]
		}
	}
	return s
}
