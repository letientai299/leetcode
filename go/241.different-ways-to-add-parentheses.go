package main

/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 *
 * https://leetcode.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (52.42%)
 * Total Accepted:    113.9K
 * Total Submissions: 200.7K
 * Testcase Example:  '"2-1-1"'
 *
 * Given a string of numbers and operators, return all possible results from
 * computing all the different possible ways to group numbers and operators.
 * The valid operators are +, - and *.
 *
 * Example 1:
 *
 *
 * Input: "2-1-1"
 * Output: [0, 2]
 * Explanation:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * Example 2:
 *
 *
 * Input: "2*3-4*5"
 * Output: [-34, -14, -10, -10, 10]
 * Explanation:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 */
func diffWaysToCompute(input string) []int {
	var nums []int
	var ops []string
	v := 0
	for _, c := range input {
		if c >= '0' && c <= '9' {
			v = 10*v + int(c-'0')
			continue
		}

		nums = append(nums, v)
		v = 0
		ops = append(ops, string(c))
	}

	nums = append(nums, v)

	calc := func(a, b int, op string) int {
		switch op {
		case "+":
			return a + b
		case "-":
			return a - b
		default:
			return a * b
		}
	}

	var compute func(ns []int, ops []string) []int

	compute = func(ns []int, ops []string) []int {
		if len(ns) == 0 {
			return nil
		}
		if len(ns) == 1 {
			return ns
		}

		if len(ns) == 2 {
			return []int{calc(ns[0], ns[1], ops[0])}
		}

		var res []int
		for i := 0; i < len(ops); i++ {
			op := ops[i]
			left := compute(ns[:i+1], ops[:i])
			right := compute(ns[i+1:], ops[i+1:])
			for _, l := range left {
				for _, r := range right {
					v := calc(l, r, op)
					res = append(res, v)
				}
			}
		}

		return res
	}

	return compute(nums, ops)
}
