package main

import "math/rand"

/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 *
 * https://leetcode.com/problems/linked-list-random-node/description/
 *
 * algorithms
 * Medium (50.59%)
 * Total Accepted:    98.4K
 * Total Submissions: 183.4K
 * Testcase Example:  '["Solution","getRandom","getRandom","getRandom","getRandom","getRandom"]\n' +
  '[[[1,2,3]],[],[],[],[],[]]'
 *
 * Given a singly linked list, return a random node's value from the linked
 * list. Each node must have the same probability of being chosen.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["Solution", "getRandom", "getRandom", "getRandom", "getRandom",
 * "getRandom"]
 * [[[1, 2, 3]], [], [], [], [], []]
 * Output
 * [null, 1, 3, 2, 2, 3]
 *
 * Explanation
 * Solution solution = new Solution([1, 2, 3]);
 * solution.getRandom(); // return 1
 * solution.getRandom(); // return 3
 * solution.getRandom(); // return 2
 * solution.getRandom(); // return 2
 * solution.getRandom(); // return 3
 * // getRandom() should return either 1, 2, or 3 randomly. Each element should
 * have equal probability of returning.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the linked list will be in the range [1, 10^4]
 * -10^4 <= Node.val <= 10^4
 * At most 10^4 calls will be made to getRandom.
 *
 *
 *
 * Follow up:
 *
 *
 * What if the linked list is extremely large and its length is unknown to
 * you?
 * Could you solve this efficiently without using extra space?
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
type Solution struct {
	m []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	s := Solution{}
	n := head
	for n != nil {
		s.m = append(s.m, n.Val)
		n = n.Next
	}
	return s
}

/** Returns a random node's value. */
func (s *Solution) GetRandom() int {
	i := rand.Int31() % int32(len(s.m))
	return s.m[i]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
