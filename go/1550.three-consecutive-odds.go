package main

func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	for i := 1; i < n-1; i++ {
		if arr[i]%2 == 1 && arr[i-1]%2 == 1 && arr[i+1]%2 == 1 {
			return true
		}
	}
	return false
}
