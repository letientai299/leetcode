package main

func mostVisited(n int, rounds []int) []int {
	sectors := make([]int, n)
	prev := rounds[0]
	for i := 1; i < len(rounds); i++ {
		for j := prev; j != rounds[i]; {
			sectors[j-1]++
			j = j%n + 1
		}
		prev = rounds[i]
		sectors[prev-1]++
		prev = prev%n + 1
	}

	max := sectors[0]
	res := []int{1}
	for i := 1; i < n; i++ {
		if sectors[i] > max {
			res = []int{i + 1}
			max = sectors[i]
			continue
		}

		if sectors[i] == max {
			res = append(res, i+1)
		}
	}

	return res
}
