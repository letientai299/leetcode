package main

/*
 * @lc app=leetcode id=477 lang=golang
 *
 * [477] Total Hamming Distance
 *
 * https://leetcode.com/problems/total-hamming-distance/description/
 *
 * algorithms
 * Medium (49.98%)
 * Total Accepted:    68.8K
 * Total Submissions: 135.9K
 * Testcase Example:  '[4,14,2]'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Now your job is to find the total Hamming distance between all pairs of the
 * given numbers.
 *
 *
 * Example:
 *
 * Input: 4, 14, 2
 *
 * Output: 6
 *
 * Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is
 * 0010 (just
 * showing the four bits relevant in this case). So the answer will be:
 * HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2
 * + 2 + 2 = 6.
 *
 *
 *
 * Note:
 *
 * Elements of the given array are in the range of 0  to 10^9
 * Length of the array will not exceed 10^4.
 *
 *
 */
func totalHammingDistance(nums []int) int {
	r := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		on := 0
		for j, v := range nums {
			on += v & 1
			nums[j] >>= 1
		}
		r += on * (n - on)
	}
	return r
}
