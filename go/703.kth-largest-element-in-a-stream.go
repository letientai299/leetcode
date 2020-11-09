package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (48.36%)
 * Total Accepted:    104.8K
 * Total Submissions: 208.9K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the k^th largest element in a stream. Note that it is
 * the k^th largest element in the sorted order, not the k^th distinct
 * element.
 *
 * Implement KthLargest class:
 *
 *
 * KthLargest(int k, int[] nums) Initializes the object with the integer k and
 * the stream of integers nums.
 * int add(int val) Returns the element representing the k^th largest element
 * in the stream.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 * Output
 * [null, 4, 5, 5, 8, 8]
 *
 * Explanation
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3);   // return 4
 * kthLargest.add(5);   // return 5
 * kthLargest.add(10);  // return 5
 * kthLargest.add(9);   // return 8
 * kthLargest.add(4);   // return 8
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= 10^4
 * 0 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * -10^4 <= val <= 10^4
 * At most 10^4 calls will be made to add.
 * It is guaranteed that there will be at least k elements in the array when
 * you search for the k^th element.
 *
 *
*/

type KthLargest struct {
	arr []int
	k   int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	c := KthLargest{
		arr: make([]int, 0, len(nums)*4),
		k:   k,
	}
	c.arr = append(c.arr, nums...)
	return c
}

func (c *KthLargest) Add(val int) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(c.arr)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if c.arr[h] < val {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}

	if i >= len(c.arr) {
		c.arr = append(c.arr, val)
		return c.kth()
	}

	for ; i < len(c.arr); i++ {
		val, c.arr[i] = c.arr[i], val
	}
	c.arr = append(c.arr, val)
	return c.kth()
}

func (c *KthLargest) kth() int {
	return c.arr[len(c.arr)-c.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
