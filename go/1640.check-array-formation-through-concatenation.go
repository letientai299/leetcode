package main

func canFormArray(arr []int, pieces [][]int) bool {
	used := make([]bool, len(arr))
	for _, ns := range pieces {
		i := 0
		for ; i < len(arr); i++ {
			if arr[i] == ns[0] {
				break
			}
		}

		if len(arr)-i < len(ns) {
			return false
		}

		for j, x := range ns {
			if arr[i+j] != x || used[i+j] {
				return false
			}
			used[i+j] = true
		}
	}
	return true
}
