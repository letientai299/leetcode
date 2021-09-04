package main

/*
 * @lc app=leetcode id=1019 lang=golang
 *
 * [1019] Next Greater Node In Linked List
 *
 * https://leetcode.com/problems/next-greater-node-in-linked-list/description/
 *
 * algorithms
 * Medium (56.60%)
 * Total Accepted:    61.6K
 * Total Submissions: 105.8K
 * Testcase Example:  '[2,1,5]'
 *
 * We are given a linked list with head as the first node.  Let's number the
 * nodes in the list: node_1, node_2, node_3, ... etc.
 *
 * Each node may have a next larger value: for node_i, next_larger(node_i) is
 * the node_j.val such that j > i, node_j.val > node_i.val, and j is the
 * smallest possible choice.  If such a j does not exist, the next larger value
 * is 0.
 *
 * Return an array of integers answer, where answer[i] =
 * next_larger(node_{i+1}).
 *
 * Note that in the example inputs (not outputs) below, arrays such as [2,1,5]
 * represent the serialization of a linked list with a head node value of 2,
 * second node value of 1, and third node value of 5.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1,5]
 * Output: [5,5,0]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,7,4,3,5]
 * Output: [7,0,5,5,0]
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1,7,5,1,9,2,5,1]
 * Output: [7,9,9,9,0,5,0,0]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= node.val <= 10^9 for each node in the linked list.
 * The given list has length in the range [0, 10000].
 *
 *
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// NOTE: solution is straightforward, speed improvement come from knowing the
// size of the list and pre-allocate necessary memory.
func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var nums []int
	for n := head; n != nil; n = n.Next {
		nums = append(nums, n.Val)
	}

	res := make([]int, len(nums))
	// pre allocate the index and self-manage the length for max speed
	index := make([]int, len(nums))
	indexLen := 0

	for h := 0; h < len(nums); h++ {
		v := nums[h]

		i := indexLen - 1
		for ; i >= 0; i-- {
			if nums[index[i]] >= v {
				break
			}

			res[index[i]] = v
		}

		indexLen = i + 1
		index[indexLen] = h
		indexLen++
	}

	return res
}
