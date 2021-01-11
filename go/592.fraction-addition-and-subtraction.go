package main

import "fmt"

/*
 * @lc app=leetcode id=592 lang=golang
 *
 * [592] Fraction Addition and Subtraction
 *
 * https://leetcode.com/problems/fraction-addition-and-subtraction/description/
 *
 * algorithms
 * Medium (47.72%)
 * Total Accepted:    20.2K
 * Total Submissions: 40.5K
 * Testcase Example:  '"-1/2+1/2"'
 *
 * Given a string representing an expression of fraction addition and
 * subtraction, you need to return the calculation result in string format. The
 * final result should be irreducible fraction. If your final result is an
 * integer, say 2, you need to change it to the format of fraction that has
 * denominator 1. So in this case, 2 should be converted to 2/1.
 *
 * Example 1:
 *
 * Input:"-1/2+1/2"
 * Output: "0/1"
 *
 *
 *
 * Example 2:
 *
 * Input:"-1/2+1/2+1/3"
 * Output: "1/3"
 *
 *
 *
 * Example 3:
 *
 * Input:"1/3-1/2"
 * Output: "-1/6"
 *
 *
 *
 * Example 4:
 *
 * Input:"5/3+1/3"
 * Output: "2/1"
 *
 *
 *
 * Note:
 *
 * The input string only contains '0' to '9', '/', '+' and '-'. So does the
 * output.
 * Each fraction (input and output) has format ±numerator/denominator. If the
 * first input fraction or the output is positive, then '+' will be omitted.
 * The input only contains valid irreducible fractions, where the numerator and
 * denominator of each fraction will always be in the range [1,10]. If the
 * denominator is 1, it means this fraction is actually an integer in a
 * fraction format defined above.
 * The number of given fractions will be in the range [1,10].
 * The numerator and denominator of the final result are guaranteed to be valid
 * and in the range of 32-bit int.
 *
 *
 */
func fractionAddition(expression string) string {
	n, d, a := 0, 1, 0 // n/d is initial value, a is running numerator
	v := 0
	sig := 1
	p := '/'
	add := func(a, b int) {
		n, d = n*b+d*a, d*b
		if x := gcdEuler(n, d); x != 1 {
			n, d = n/x, d/x
		}
	}
	for _, c := range expression {
		switch c {
		case '/':
			a = sig * v
			v = 0
			sig = 1
		case '+':
			sig = 1
			add(a*sig, v)
			v = 0
		case '-':
			if p != '/' {
				add(a*sig, v)
				v = 0
			}
			sig = -1
		default:
			v = 10*v + int(c-'0')
		}
		p = c
	}

	add(a*sig, v)

	if d < 0 {
		n, d = -n, -d
	}
	return fmt.Sprintf("%d/%d", n, d)
}
