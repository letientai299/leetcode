package main

/*
* @lc app=leetcode id=303 lang=golang
*
* [303] Range Sum Query - Immutable
*
* https://leetcode.com/problems/range-sum-query-immutable/description/
*
* algorithms
* Easy (37.02%)
* Total Accepted:    130.7K
* Total Submissions: 353K
* Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
*
* Given an integer array nums, find the sum of the elements between indices i
* and j (i ≤ j), inclusive.
*
* Example:
*
* Given nums = [-2, 0, 3, -5, 2, -1]
*
* sumRange(0, 2) -> 1
* sumRange(2, 5) -> -1
* sumRange(0, 5) -> -3
*
*
*
* Note:
*
* You may assume that the array does not change.
* There are many calls to sumRange function.
*
*
 */
type NumArray struct {
	ns []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return NumArray{
		ns: nums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if len(this.ns) == 0 {
		return 0
	}
	if i == 0 {
		return this.ns[j]
	}

	return this.ns[j] - this.ns[i-1]
}

/**
* Your NumArray object will be instantiated and called as such:
* obj := Constructor(nums);
* param_1 := obj.SumRange(i,j);
 */
