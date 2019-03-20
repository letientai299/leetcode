package main

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (51.08%)
 * Total Accepted:    298.4K
 * Total Submissions: 584.2K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (i.e., buy one and sell one share of the stock
 * multiple times).
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit
 * = 5-1 = 4.
 * Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 =
 * 3.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// remove location where price are unchanged
	x, y := 0, 1
	for x < n && y < n {
		if prices[x] != prices[y] {
			prices[x+1] = prices[y]
			y++
			x++
			continue
		}

		for y < n && prices[x] == prices[y] {
			y++
		}
	}

	prices = prices[0 : x+1]
	n = len(prices)

	if n <= 1 {
		return 0
	}

	if n == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		}
	}

	// store the price points when it changes its trend
	var points []int

	// consider first points is an lower points if it less than the second prices
	if prices[0] <= prices[1] {
		// denote lower points as negative value, as we would spend money to buy
		// stock at this point
		points = append(points, -prices[0])
	}

	lastPrice := prices[0]
	for i := 1; i < n-1; i++ {
		currentPrice := prices[i]
		nextPrice := prices[i+1]

		// if this is a top point, where price is at peak
		if lastPrice < currentPrice && currentPrice > nextPrice {
			points = append(points, currentPrice) // local price peak
		}

		if lastPrice > currentPrice && currentPrice < nextPrice {
			points = append(points, -currentPrice) // local bottom
		}

		// if this is a bottom point, where price is at its lowest
		lastPrice = currentPrice
	}

	// consider last point is an upper point
	if prices[n-2] <= prices[n-1] {
		points = append(points, prices[n-1])
	}

	sum := 0
	for _, x := range points {
		sum += x
	}
	return sum
}
