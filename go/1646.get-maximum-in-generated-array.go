package main

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	max := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
			continue
		}

		arr[i] = arr[i/2] + arr[i/2+1]
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
