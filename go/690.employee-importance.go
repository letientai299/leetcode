package main

/*
 * @lc app=leetcode id=690 lang=golang
 *
 * [690] Employee Importance
 *
 * https://leetcode.com/problems/employee-importance/description/
 *
 * algorithms
 * Easy (55.99%)
 * Total Accepted:    88.4K
 * Total Submissions: 152.7K
 * Testcase Example:  '[[1,2,[2]], [2,3,[]]]\n2'
 *
 * You are given a data structure of employee information, which includes the
 * employee's unique id, their importance value and their direct subordinates'
 * id.
 *
 * For example, employee 1 is the leader of employee 2, and employee 2 is the
 * leader of employee 3. They have importance value 15, 10 and 5, respectively.
 * Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has
 * [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3
 * is also a subordinate of employee 1, the relationship is not direct.
 *
 * Now given the employee information of a company, and an employee id, you
 * need to return the total importance value of this employee and all their
 * subordinates.
 *
 * Example 1:
 *
 *
 * Input: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
 * Output: 11
 * Explanation:
 * Employee 1 has importance value 5, and he has two direct subordinates:
 * employee 2 and employee 3. They both have importance value 3. So the total
 * importance value of employee 1 is 5 + 3 + 3 = 11.
 *
 *
 *
 *
 * Note:
 *
 *
 * One employee has at most one direct leader and may have several
 * subordinates.
 * The maximum number of employees won't exceed 2000.
 *
 *
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	em := make(map[int]*Employee, len(employees))

	arr := employees
	var unknown []int
	for len(arr) != 0 {
		next := unknown
		for _, e := range arr {
			em[e.Id] = e
			next = append(next, e.Subordinates...)
		}
		arr = nil
		for _, id := range next {
			e, ok := em[id]
			if ok {
				arr = append(arr, e)
			} else {
				unknown = append(unknown, id)
			}
		}
	}

	e, ok := em[id]
	if !ok {
		return 0
	}
	res := e.Importance
	subs := e.Subordinates

	for len(subs) != 0 {
		var next []int
		for _, id := range subs {
			e, ok := em[id]
			if ok {
				res += e.Importance
				next = append(next, e.Subordinates...)
			}
		}
		subs = next
	}

	return res
}
