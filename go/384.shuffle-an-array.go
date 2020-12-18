package main

import "math/rand"

/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 *
 * https://leetcode.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (51.29%)
 * Total Accepted:    152.1K
 * Total Submissions: 283.9K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * Given an integer array nums, design an algorithm to randomly shuffle the
 * array.
 *
 * Implement the Solution class:
 *
 *
 * Solution(int[] nums) Initializes the object with the integer array nums.
 * int[] reset() Resets the array to its original configuration and returns
 * it.
 * int[] shuffle() Returns a random shuffling of the array.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["Solution", "shuffle", "reset", "shuffle"]
 * [[[1, 2, 3]], [], [], []]
 * Output
 * [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
 *
 * Explanation
 * Solution solution = new Solution([1, 2, 3]);
 * solution.shuffle();    // Shuffle the array [1,2,3] and return its result.
 * Any permutation of [1,2,3] must be equally likely to be returned. Example:
 * return [3, 1, 2]
 * solution.reset();      // Resets the array back to its original
 * configuration [1,2,3]. Return [1, 2, 3]
 * solution.shuffle();    // Returns the random shuffling of array [1,2,3].
 * Example: return [1, 3, 2]
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * -10^6 <= nums[i] <= 10^6
 * All the elements of nums are unique.
 * At most 5 * 10^4 calls will be made to reset and shuffle.
 *
 *
 */
type Solution struct {
	arr []int
	b   []int
}

func Constructor384(nums []int) Solution {
	b := make([]int, len(nums))
	copy(b, nums)
	return Solution{
		arr: nums,
		b:   b,
	}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.arr
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	rand.Shuffle(len(s.arr), func(i, j int) {
		s.b[i], s.b[j] = s.b[j], s.b[i]
	})
	return s.b
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
