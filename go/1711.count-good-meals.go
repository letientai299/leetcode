package main

// medium
// 1711. Count Good Meals
func countPairs(deliciousness []int) int {
	m := make(map[int]int)
	for _, d := range deliciousness {
		m[d]++
	}

	const mod = 1e9 + 7

	r := 0
	for k, v := range m {
		for i := 0; i <= 21; i++ {
			base := 1 << i
			if base < k {
				continue
			}

			if base > k*2 {
				break
			}

			if base == k*2 {
				r = (r + v*(v-1)/2) % mod
			} else {
				u := m[base-k]
				r = (r + v*u) % mod
			}
		}
	}

	return r
}
