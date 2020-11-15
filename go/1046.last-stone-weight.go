package main

import "container/heap"

/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 *
 * https://leetcode.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (62.41%)
 * Total Accepted:    149.6K
 * Total Submissions: 239.7K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * We have a collection of stones, each stone has a positive integer weight.
 *
 * Each turn, we choose the two heaviest stones and smash them together.
 * Suppose the stones have weights x and y with x <= y.  The result of this
 * smash is:
 *
 *
 * If x == y, both stones are totally destroyed;
 * If x != y, the stone of weight x is totally destroyed, and the stone of
 * weight y has new weight y-x.
 *
 *
 * At the end, there is at most 1 stone left.  Return the weight of this stone
 * (or 0 if there are no stones left.)
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the
 * value of last stone.
 *
 *
 *
 * Note:
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 *
 *
 */
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	h := &intHeap{store: stones}
	heap.Init(h)
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		if a != b {
			heap.Push(h, a-b)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h.Pop().(int)
}

type intHeap struct {
	store []int
}

func (h *intHeap) Len() int {
	return len(h.store)
}

func (h *intHeap) Less(i, j int) bool {
	return h.store[i] > h.store[j]
}

func (h *intHeap) Swap(i, j int) {
	h.store[i], h.store[j] = h.store[j], h.store[i]
}

func (h *intHeap) Push(x interface{}) {
	h.store = append(h.store, x.(int))
}

func (h *intHeap) Pop() interface{} {
	n := len(h.store)
	if n == 0 {
		return nil
	}
	x := h.store[n-1]
	h.store = h.store[:n-1]
	return x
}
