package main

import "container/heap"

// Merge k Sorted Lists
//
// Hard
//
// https://leetcode.com/problems/merge-k-sorted-lists/
//
// You are given an array of `k` linked-lists `lists`, each linked-list is
// sorted in ascending order.
//
// _Merge all the linked-lists into one sorted linked-list and return it._
//
// **Example 1:**
//
// ```
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
//
// ```
//
// **Example 2:**
//
// ```
// Input: lists = []
// Output: []
//
// ```
//
// **Example 3:**
//
// ```
// Input: lists = [[]]
// Output: []
//
// ```
//
// **Constraints:**
//
// - `k == lists.length`
// - `0 <= k <= 10^4`
// - `0 <= lists[i].length <= 500`
// - `-10^4 <= lists[i][j] <= 10^4`
// - `lists[i]` is sorted in **ascending order**.
// - The sum of `lists[i].length` won't exceed `10^4`.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	i := 0
	for _, l := range lists{
		if l != nil {
			lists[i] = l
			i++
		}
	}
	if i == 0 {
		return nil
	}

	lists = lists[:i]

	h := ListNodeHeap(lists)
	pre := &ListNode{}
	cur := pre
	heap.Init(&h)

	for len(h) != 0 {
		node := heap.Pop(&h).(*ListNode)
		if node == nil {
			continue
		}

		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}

	return pre.Next
}

type ListNodeHeap []*ListNode

func (l ListNodeHeap) Len() int            { return len(l) }
func (l ListNodeHeap) Less(i, k int) bool  { return l[i].Val < l[k].Val }
func (l ListNodeHeap) Swap(i, j int)       { l[i], l[j] = l[j], l[i] }
func (l *ListNodeHeap) Push(x interface{}) { *l = append(*l, x.(*ListNode)) }

func (l *ListNodeHeap) Pop() interface{} {
	n := len(*l)
	v := (*l)[n-1]
	*l = (*l)[:n-1]
	return v
}
