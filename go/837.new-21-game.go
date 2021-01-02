package main

/*
 * @lc app=leetcode id=837 lang=golang
 *
 * [837] New 21 Game
 *
 * https://leetcode.com/problems/new-21-game/description/
 *
 * algorithms
 * Medium (33.18%)
 * Total Accepted:    20.4K
 * Total Submissions: 58K
 * Testcase Example:  '10\n1\n10'
 *
 * Alice plays the following game, loosely based on the card game "21".
 *
 * Alice starts with 0 points, and draws numbers while she has less than K
 * points.  During each draw, she gains an integer number of points randomly
 * from the range [1, W], where W is an integer.  Each draw is independent and
 * the outcomes have equal probabilities.
 *
 * Alice stops drawing numbers when she gets K or more points.  What is the
 * probability that she has N or less points?
 *
 * Example 1:
 *
 *
 * Input: N = 10, K = 1, W = 10
 * Output: 1.00000
 * Explanation:  Alice gets a single card, then stops.
 *
 *
 * Example 2:
 *
 *
 * Input: N = 6, K = 1, W = 10
 * Output: 0.60000
 * Explanation:  Alice gets a single card, then stops.
 * In 6 out of W = 10 possibilities, she is at or below N = 6 points.
 *
 *
 * Example 3:
 *
 *
 * Input: N = 21, K = 17, W = 10
 * Output: 0.73278
 *
 * Note:
 *
 *
 * 0 <= K <= N <= 10000
 * 1 <= W <= 10000
 * Answers will be accepted as correct if they are within 10^-5 of the correct
 * answer.
 * The judging time limit has been reduced for this question.
 *
 *
 */
// TODO (tai): correct, but TLE
func new21Game(n int, k int, w int) float64 {
	if n >= k-1+w || w == 1 || k == 0 {
		return 1
	}

	if n < k {
		return 0
	}

	if k == 1 {
		return float64(n) / float64(w)
	}

	mem := make(map[int]float64)
	mem[0] = 1
	var prob func(x int) float64
	prob = func(x int) float64 {
		if v, ok := mem[x]; ok {
			return v
		}

		s := float64(0)
		for i := 1; i <= w; i++ {
			if x < i {
				break
			}
			s += prob(x - i)
		}
		s /= float64(w)
		mem[x] = s
		return s
	}

	s := prob(k)
	for i := 1; i <= n-k; i++ {
		for j := k + i - w; j <= k-1; j++ {
			v := prob(j) / float64(w)
			s += v
		}
	}

	return s
}
