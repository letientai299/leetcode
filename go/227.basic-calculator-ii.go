package main

/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 *
 * https://leetcode.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (35.15%)
 * Total Accepted:    232.9K
 * Total Submissions: 611.8K
 * Testcase Example:  '"3+2*2"'
 *
 * Given a string s which represents an expression, evaluate this expression
 * and return its value. 
 *
 * The integer division should truncate toward zero.
 *
 *
 * Example 1:
 * Input: s = "3+2*2"
 * Output: 7
 * Example 2:
 * Input: s = " 3/2 "
 * Output: 1
 * Example 3:
 * Input: s = " 3+5 / 2 "
 * Output: 5
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s consists of integers and operators ('+', '-', '*', '/') separated by some
 * number of spaces.
 * s represents a valid expression.
 * All the integers in the expression are non-negative integers in the range
 * [0, 2^31 - 1].
 * The answer is guaranteed to fit in a 32-bit integer.
 *
 *
 */
func calculate(s string) int {
	var ns []int
	var ops []byte

	v := 0
	for _, x := range s {
		switch x {
		case '+':
			ops = append(ops, byte(x))
			ns = append(ns, v)
			v = 0
		case '-':
			ops = append(ops, byte(x))
			ns = append(ns, v)
			v = 0
		case '*':
			ops = append(ops, byte(x))
			ns = append(ns, v)
			v = 0
		case '/':
			ops = append(ops, byte(x))
			ns = append(ns, v)
			v = 0
		case ' ':
			continue
		default:
			v = 10*v + int(x-'0')
		}
	}
	ns = append(ns, v)

	for i, o := range ops {
		switch o {
		case '-':
			ns[i+1] = -ns[i+1]
		case '*':
			ns[i+1] *= ns[i]
			ns[i] *= 0
		case '/':
			ns[i+1] = ns[i] / ns[i+1]
			ns[i] *= 0
		default:
			continue
		}
	}

	v = 0
	for _, x := range ns {
		v += x
	}

	return v
}
