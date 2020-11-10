package main /*
 * @lc app=leetcode id=1010 lang=golang
 *
 * [1010] Pairs of Songs With Total Durations Divisible by 60
 *
 * https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/
 *
 * algorithms
 * Easy (46.59%)
 * Total Accepted:    44K
 * Total Submissions: 92K
 * Testcase Example:  '[30,20,150,100,40]'
 *
 * In a list of songs, the i-th song has a duration of time[i] seconds.
 *
 * Return the number of pairs of songs for which their total duration in
 * seconds is divisible by 60.  Formally, we want the number of indices i, j
 * such that i < j with (time[i] + time[j]) % 60 == 0.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [30,20,150,100,40]
 * Output: 3
 * Explanation: Three pairs have a total duration divisible by 60:
 * (time[0] = 30, time[2] = 150): total duration 180
 * (time[1] = 20, time[3] = 100): total duration 120
 * (time[1] = 20, time[4] = 40): total duration 60
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [60,60,60]
 * Output: 3
 * Explanation: All three pairs have a total duration of 120, which is
 * divisible by 60.
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= time.length <= 60000
 * 1 <= time[i] <= 500
 *
 *
 */
func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	for _, v := range time {
		m[v%60]++
	}

	p := 0

	for k, v := range m {
		r := (60 - k) % 60
		if r == k {
			p += v * (v - 1)
			continue
		}

		p += v * m[r]
	}
	return p / 2
}
