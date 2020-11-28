package main

/*
 * @lc app=leetcode id=781 lang=golang
 *
 * [781] Rabbits in Forest
 *
 * https://leetcode.com/problems/rabbits-in-forest/description/
 *
 * algorithms
 * Medium (52.76%)
 * Total Accepted:    23.2K
 * Total Submissions: 42K
 * Testcase Example:  '[1,0,1,0,0]'
 *
 * In a forest, each rabbit has some color. Some subset of rabbits (possibly
 * all of them) tell you how many other rabbits have the same color as them.
 * Those answers are placed in an array.
 *
 * Return the minimum number of rabbits that could be in the forest.
 *
 *
 * Examples:
 * Input: answers = [1, 1, 2]
 * Output: 5
 * Explanation:
 * The two rabbits that answered "1" could both be the same color, say red.
 * The rabbit than answered "2" can't be red or the answers would be
 * inconsistent.
 * Say the rabbit that answered "2" was blue.
 * Then there should be 2 other blue rabbits in the forest that didn't answer
 * into the array.
 * The smallest possible number of rabbits in the forest is therefore 5: 3 that
 * answered plus 2 that didn't.
 *
 * Input: answers = [10, 10, 10]
 * Output: 11
 *
 * Input: answers = []
 * Output: 0
 *
 *
 * Note:
 *
 *
 * answers will have length at most 1000.
 * Each answers[i] will be an integer in the range [0, 999].
 *
 *
 */
func numRabbits(answers []int) int {
	m := make(map[int]int)
	for _, v := range answers {
		m[v]++
	}

	r := 0
	for k, v := range m {
		// this line is 0ms (100%)
		r += (k + 1) * ((v + k) / (k + 1))
		// the following is 4-8ms, good trick to learn
		// r += (k + 1) * (v/(k + 1))
		// if v%(k+1) != 0 {
		//	r += k + 1
		// }
	}

	return r
}
