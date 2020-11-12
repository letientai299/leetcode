package main

func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)
	for i := 0; i < n-m; i++ {
		if i+m*k-1 >= n {
			return false
		}

		ok := true
		for x := 1; x < k && ok; x++ {
			for j := 0; j < m; j++ {
				h := j + m*x + i
				if h >= n || arr[h] != arr[i+j] {
					ok = false
					break
				}
			}
		}
		if ok {
			return true
		}
	}
	return false
}
