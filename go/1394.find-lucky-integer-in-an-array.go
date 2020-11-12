package main

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, x := range arr {
		m[x]++
	}
	max := -1
	for k, v := range m {
		if v != k {
			continue
		}

		if k > max {
			max = k
		}
	}

	return max
}
