package main

func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for i, v := range arr {
		if _, ok := m[v*2]; ok {
			return true
		}

		if v%2 == 0 {
			if _, ok := m[v/2]; ok {
				return true
			}
		}

		m[v] = i
	}

	return false
}
