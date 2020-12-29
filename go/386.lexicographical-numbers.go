package main

/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (48.59%)
 * Total Accepted:    60.4K
 * Total Submissions: 113.1K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */
func lexicalOrder(n int) []int {
	if n == 0 {
		return nil
	}

	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = lexicalOrderNext(arr[i-1], n)
	}
	return arr
}

func lexicalOrderNext(n int, cap int) int {
	n *= 10 // append 0
	if n <= cap {
		return n
	}
	n /= 10
	n++
	if n > cap {
		if n%10 == 0 {
			n /= 10
		} else {
			n /= 10
			n++
		}
		for n%10 == 0 {
			n /= 10
		}
		return n
	}

	for n%10 == 0 {
		n /= 10
	}
	return n
}
