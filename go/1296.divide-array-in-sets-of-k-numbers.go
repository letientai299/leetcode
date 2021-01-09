package main

// medium
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
func isPossibleDivide(nums []int, k int) bool {
	return isNStraightHand(nums, k)
}
