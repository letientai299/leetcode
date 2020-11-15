package main

func decrypt(code []int, k int) []int {
	n := len(code)
	sum := make([]int, n)
	if k == 0 {
		return sum
	}

	sum[0] = code[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + code[i]
	}

	res := make([]int, n)
	for i, _ := range code {
		var start, end int
		if k > 0 {
			end = (i + k) % n
			start = (i + 1) % n
		} else {
			end = (i - 1 + n) % n
			start = (i + k + n) % n
		}

		if start <= end {
			res[i] = sum[end] - sum[start] + code[start]
		} else {
			res[i] = sum[n-1] - sum[start] + sum[end] + code[start]
		}
	}
	return res
}
