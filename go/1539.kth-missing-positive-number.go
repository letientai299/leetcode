package main

func findKthPositive(arr []int, k int) int {
	if len(arr) == 0 {
		return k
	}

	prev := arr[0]
	k -= prev - 1

	for i := 1; i < len(arr); i++ {
		if k <= 0 {
			break
		}
		k -= arr[i] - prev - 1
		prev = arr[i]
	}

	if k <= 0 {
		return prev + k - 1
	}

	return prev + k
}
