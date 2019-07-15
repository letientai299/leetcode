package main

func max(a int, arr ...int) int {
	for _, v := range arr {
		if a < v {
			a = v
		}
	}

	return a
}
