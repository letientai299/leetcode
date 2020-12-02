package main

import (
	"math"
	"sort"
)

// medium

// TODO (tai): slow, 48.39%, 168ms
func maxProfit(inventory []int, orders int) int {
	m := make(map[int]int)
	k := 0
	for _, v := range inventory {
		m[v]++
		if v > k {
			k = v
		}
	}

	r := 0

	n := len(inventory)
	if float64(k) > math.Log2(float64(n))*float64(n) {
		sort.Ints(inventory)
		factor := 1
		for i := n - 1; i > 0; i-- {
			x := inventory[i]
			sell := x - inventory[i-1]
			if sell == 0 {
				factor++
				continue
			}

			if orders <= factor*sell {
				sell = orders / factor
			}

			orders -= sell * factor
			r = (r + factor*(x+(x-sell+1))*sell/2) % (1e9 + 7)

			if sell < x-inventory[i-1] {
				x -= sell
				for orders > 0 {
					for j := 0; j < factor && orders > 0; j++ {
						r = (r + x) % (1e9 + 7)
						orders--
					}
					x--
				}
				return r
			}

			factor++
		}

		x := inventory[0]
		sell := x
		if orders <= factor*x {
			sell = orders / factor
		}
		orders -= sell * factor
		r = (r + factor*(x+(x-sell+1))*sell/2) % (1e9 + 7)
		x -= sell

		for orders > 0 {
			for j := 0; j < factor && orders > 0; j++ {
				r = (r + x) % (1e9 + 7)
				orders--
			}
			x--
		}
		return r
	}

	for orders > 0 {
		for m[k] == 0 {
			k--
		}

		if orders > m[k] {
			orders -= m[k]
			r = (r + k*m[k]) % (1e9 + 7)
			m[k-1] += m[k]
			k--
		} else {
			r = (r + k*orders) % (1e9 + 7)
			orders = 0
		}
	}
	return r
}
