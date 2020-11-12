package main

func numSpecial(mat [][]int) int {
	var potentials []int
	for _, r := range mat {
		c := -1

		for i, v := range r {
			if v == 1 {
				if c >= 0 {
					c = -1
					break
				}
				c = i
			}
		}

		potentials = append(potentials, c)
	}

	res := 0
	for r, c := range potentials {
		if c == -1 {
			continue
		}

		yes := true
		for i := 0; i < len(mat); i++ {
			if i == r {
				continue
			}

			if mat[i][c] == 1 {
				yes = false
				break
			}
		}
		if yes {
			res++
		}
	}
	return res
}
