package main

/*
 * @lc app=leetcode id=339 lang=golang
 *
 * [339] Nested List Weight Sum
 *
 * https://leetcode.com/problems/nested-list-weight-sum/description/
 *
 * algorithms
 * Easy (70.66%)
 * Total Accepted:    101.2K
 * Total Submissions: 135.3K
 * Testcase Example:  '[[1,1],2,[1,1]]'
 *
 * Given a nested list of integers, return the sum of all integers in the list
 * weighted by their depth.
 *
 * Each element is either an integer, or a list -- whose elements may also be
 * integers or other lists.
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,1],2,[1,1]]
 * Output: 10
 * Explanation: Four 1's at depth 2, one 2 at depth 1.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,[4,[6]]]
 * Output: 27
 * Explanation: One 1 at depth 1, one 4 at depth 2, and one 6 at depth 3; 1 +
 * 4*2 + 6*3 = 27.
 *
 *
 *
 */
func depthSum(nestedList []*NestedInteger) int {
	var sum func(ns []*NestedInteger, d int) int
	sum = func(ns []*NestedInteger, d int) int {
		s := 0
		for _, n := range ns {
			if n.IsInteger() {
				s += n.GetInteger() * d
				continue
			}

			s += sum(n.GetList(), d+1)
		}
		return s
	}

	return sum(nestedList, 1)
}

// Stupid "interface" from Leetcode, copy here to make the code compilable.
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }
