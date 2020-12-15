package main

/*
 * @lc app=leetcode id=385 lang=golang
 *
 * [385] Mini Parser
 *
 * https://leetcode.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (32.60%)
 * Total Accepted:    40.7K
 * Total Submissions: 119.1K
 * Testcase Example:  '"324"'
 *
 * Given a nested list of integers represented as a string, implement a parser
 * to deserialize it.
 *
 * Each element is either an integer, or a list -- whose elements may also be
 * integers or other lists.
 *
 * Note: You may assume that the string is well-formed:
 *
 *
 * String is non-empty.
 * String does not contain white spaces.
 * String contains only digits 0-9, [, - ,, ].
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Given s = "324",
 *
 * You should return a NestedInteger object which contains a single integer
 * 324.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Given s = "[123,[456,[789]]]",
 *
 * Return a NestedInteger object containing a nested list with 2 elements:
 *
 * 1. An integer containing value 123.
 * 2. A nested list containing two elements:
 * ⁠   i.  An integer containing value 456.
 * ⁠   ii. A nested list with one element:
 * ⁠        a. An integer containing value 789.
 *
 *
 *
 *
 */
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	// pointer to current list we add items to
	var list *NestedInteger
	// stack of parent lists, when we encounter inner array
	var stack []*NestedInteger

	// num is a number we calculate
	// numType is status of the number:
	// 0 - no number
	// 1 - there is a number
	// -1 - there was a negative sign before a number
	num, numType := 0, 0
	for _, c := range s {
		switch c {
		case '[':
			// append current list to stack and assign new pointer to `list`
			if list != nil {
				stack = append(stack, list)
			}
			list = &NestedInteger{}
		case ']':
			// if previously was calculated number, add it to list
			if numType > 0 {
				elem := NestedInteger{}
				elem.SetInteger(num)
				list.Add(elem)
				num, numType = 0, 0
			}
			// Take out parent list, add current list to it's parent and assign `list` = it's parent
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Add(*list)
				list = parent
				stack = stack[:len(stack)-1]
			}
		case ',':
			// if previously was calculated number, add it to list
			if list == nil {
				list = &NestedInteger{}
			}
			if numType > 0 {
				elem := NestedInteger{}
				elem.SetInteger(num)
				list.Add(elem)
				num, numType = 0, 0
			}
		case '-':
			numType = -1
		default:
			if numType == 0 {
				numType = 1
			}
			if num < 0 {
				num = num*10 - int(c-'0')
			} else {
				num = num*10 + int(c-'0')
			}
			// revert sign only once to have less `if's` further when we add element to list
			if numType < 0 {
				num = -num
				numType = 1
			}
		}
	}
	// if previously was calculated number but not added then add it to the list
	if numType > 0 {
		if list == nil {
			list = &NestedInteger{}
			list.SetInteger(num)
		} else {
			elem := NestedInteger{}
			elem.SetInteger(num)
			list.Add(elem)
		}
	}

	return list
}
