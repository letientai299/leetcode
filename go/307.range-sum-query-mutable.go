package main

/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 *
 * https://leetcode.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (31.17%)
 * Total Accepted:    123.5K
 * Total Submissions: 341.1K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * The update(i, val) function modifies nums by updating the element at index i
 * to val.
 *
 * Example:
 *
 *
 * Given nums = [1, 3, 5]
 *
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 *
 *
 *
 * Constraints:
 *
 *
 * The array is only modifiable by the update function.
 * You may assume the number of calls to update and sumRange function is
 * distributed evenly.
 * 0 <= i <= j <= nums.length - 1
 *
 *
 */

// TODO (tai): slow, 6.12%
type NumArray struct {
	arr []int
	sum []int
}

func Constructor(nums []int) NumArray {
	a := NumArray{
		arr: nums,
		sum: make([]int, len(nums)+1),
	}

	v := 0
	for i, x := range nums {
		a.sum[i+1] = v + x
		v += x
	}

	return a
}

func (a *NumArray) Update(i int, val int) {
	d := a.arr[i] - val
	a.arr[i] = val
	for j := i; j < len(a.arr); j++ {
		a.sum[j+1] -= d
	}
}

func (a *NumArray) SumRange(i int, j int) int {
	return a.sum[j+1] - a.sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
