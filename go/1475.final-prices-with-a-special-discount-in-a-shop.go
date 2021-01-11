package main

func finalPrices(prices []int) []int {
	for i, x := range prices {
		for j := i + 1; j < len(prices); j++ {
			if x >= prices[j] {
				prices[i] = x - prices[j]
				break
			}
		}
	}
	return prices
}
