package main

/*
 * @lc app=leetcode id=170 lang=golang
 *
 * [170] Two Sum III - Data structure design
 *
 * https://leetcode.com/problems/two-sum-iii-data-structure-design/description/
 *
 * algorithms
 * Easy (31.80%)
 * Total Accepted:    68.9K
 * Total Submissions: 215K
 * Testcase Example:  '["TwoSum","add","add","add","find","find"]\n[[],[1],[3],[5],[4],[7]]'
 *
 * Design and implement a TwoSum class. It should support the following
 * operations: add and find.
 *
 * add - Add the number to an internal data structure.
 * find - Find if there exists any pair of numbers which sum is equal to the
 * value.
 *
 * Example 1:
 *
 *
 * add(1); add(3); add(5);
 * find(4) -> true
 * find(7) -> false
 *
 *
 * Example 2:
 *
 *
 * add(3); add(1); add(2);
 * find(3) -> true
 * find(6) -> false
 *
 */
type TwoSum struct {
	m map[int]int
	l []int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{m: make(map[int]int)}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.m[number]++
	this.l = append(this.l, number)
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for _, n := range this.l {
		n2 := value - n
		if (n == n2 && this.m[n] > 1) || (n != n2 && this.m[n2] > 0) {
			return true
		}
	}
	return false
}
