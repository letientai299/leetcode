package main

import "math"

/*
 * @lc app=leetcode id=319 lang=golang
 *
 * [319] Bulb Switcher
 *
 * https://leetcode.com/problems/bulb-switcher/description/
 *
 * algorithms
 * Medium (44.54%)
 * Total Accepted:    67.3K
 * Total Submissions: 150.7K
 * Testcase Example:  '3'
 *
 * There are n bulbs that are initially off. You first turn on all the bulbs.
 * Then, you turn off every second bulb. On the third round, you toggle every
 * third bulb (turning on if it's off or turning off if it's on). For the i-th
 * round, you toggle every i bulb. For the n-th round, you only toggle the last
 * bulb. Find how many bulbs are on after n rounds.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 1
 * Explanation:
 * At first, the three bulbs are [off, off, off].
 * After first round, the three bulbs are [on, on, on].
 * After second round, the three bulbs are [on, off, on].
 * After third round, the three bulbs are [on, off, off].
 *
 * So you should return 1, because there is only one bulb is on.
 *
 *
 */
func bulbSwitch(n int) int {
	//bulbs := make([]bool, n)
	//for i := 1; i <= n; i++ {
	//	for j := i; j <= n; j += i {
	//		bulbs[j-1] = !bulbs[j-1]
	//	}
	//}
	//
	//count := 0
	//for i := 0; i < n; i++ {
	//	if bulbs[i] {
	//		count++
	//	}
	//}

	// I can see the pattern by testing, but can't derive the logic of perfect
	// square from that within 5m.
	// https://leetcode.com/problems/bulb-switcher/discuss/402164/Full-Description-and-Justification
	return int(math.Floor(math.Sqrt(float64(n))))
}
