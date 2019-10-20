package main

/*
 * @lc app=leetcode id=754 lang=golang
 *
 * [754] Reach a Number
 *
 * https://leetcode.com/problems/reach-a-number/description/
 *
 * algorithms
 * Easy (33.51%)
 * Total Accepted:    13.2K
 * Total Submissions: 39.3K
 * Testcase Example:  '1'
 *
 *
 * You are standing at position 0 on an infinite number line.  There is a goal
 * at position target.
 *
 * On each move, you can either go left or right.  During the n-th move
 * (starting from 1), you take n steps.
 *
 * Return the minimum number of steps required to reach the destination.
 *
 *
 * Example 1:
 *
 * Input: target = 3
 * Output: 2
 * Explanation:
 * On the first move we step from 0 to 1.
 * On the second step we step from 1 to 3.
 *
 *
 *
 * Example 2:
 *
 * Input: target = 2
 * Output: 3
 * Explanation:
 * On the first move we step from 0 to 1.
 * On the second move we step  from 1 to -1.
 * On the third move we step from -1 to 2.
 *
 *
 *
 * Note:
 * target will be a non-zero integer in the range [-10^9, 10^9].
 *
 */

// TODO (jl): try to solve this again, I'm not quite sure about all the math.
func reachNumber(target int) int {
	// the solution for negative target is same with positive one
	if target < 0 {
		target = -target
	}

	k := 0
	sum := 0

	// there's some ideas to get rid of this for loop, but I don't copy them,
	// because I don't understand them
	for sum < target || (sum-target)%2 != 0 {
		k++
		sum += k
	}

	return k
}
