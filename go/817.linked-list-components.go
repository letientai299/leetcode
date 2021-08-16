package main

// Linked List Components
//
// Medium
//
// https://leetcode.com/problems/linked-list-components/
//
// You are given the `head` of a linked list containing unique integer values
// and an integer array `nums` that is a subset of the linked list values.
//
// Return _the number of connected components in_ `nums` _where two values are
// connected if they appear **consecutively** in the linked list_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom1.jpg)
//
// ```
// Input: head = [0,1,2,3], nums = [0,1,3]
// Output: 2
// Explanation: 0 and 1 are connected, so [0, 1] and [3] are the two connected
// components.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom2.jpg)
//
// ```
// Input: head = [0,1,2,3,4], nums = [0,3,1,4]
// Output: 2
// Explanation: 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3,
// 4] are the two connected components.
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the linked list is `n`.
// - `1 <= n <= 104`
// - `0 <= Node.val < n`
// - All the values `Node.val` are **unique**.
// - `1 <= nums.length <= n`
// - `0 <= nums[i] <= n`
// - All the values of `nums` are **unique**.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	vals := make(map[int]struct{})
	for _, v := range nums {
		vals[v] = struct{}{}
	}

	r := 0
	pre := -2
	for i, node := 0, head; node != nil; node = node.Next {
		v := node.Val
		if _, care := vals[v]; care {
			if pre != i-1 {
				r++
			}
			pre = i
		}
		i++
	}

	return r
}

// Good for interview
