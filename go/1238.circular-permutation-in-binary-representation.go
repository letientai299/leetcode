package main

/*
 * @lc app=leetcode id=1238 lang=golang
 *
 * [1238] Circular Permutation in Binary Representation
 *
 * https://leetcode.com/problems/circular-permutation-in-binary-representation/description/
 *
 * algorithms
 * Medium (58.81%)
 * Total Accepted:    8.3K
 * Total Submissions: 12.6K
 * Testcase Example:  '2\n3'
 *
 * Given 2 integers n and start. Your task is return any permutation p of
 * (0,1,2.....,2^n -1) such that :
 *
 *
 * p[0] = start
 * p[i] and p[i+1] differ by only one bit in their binary representation.
 * p[0] and p[2^n -1] must also differ by only one bit in their binary
 * representation.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2, start = 3
 * Output: [3,2,0,1]
 * Explanation: The binary representation of the permutation is (11,10,00,01).
 * All the adjacent element differ by one bit. Another valid permutation is
 * [3,1,0,2]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3, start = 2
 * Output: [2,6,7,5,4,0,1,3]
 * Explanation: The binary representation of the permutation is
 * (010,110,111,101,100,000,001,011).
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 16
 * 0 <= start < 2 ^ n
 *
 */
func circularPermutation(n int, start int) []int {
	arr := make([]int, 1<<n)
	arr[0] = 0
	arr[1] = 1

	var split int
	if start == 0 {
		split = 0
	} else if start == 1 {
		split = 1
	}

	for i := 1; i < n; i++ {
		l := 1 << i
		up := 2*l - 1
		for j := l - 1; j >= 0; j-- {
			arr[up-j] = l + arr[j]
			if arr[up-j] == start {
				split = up - j
			}
		}
	}

	//format := fmt.Sprintf("%%0%db\n", n)
	//for _, v := range arr {
	//	fmt.Printf(format, v)
	//}

	if split == 0 {
		return arr
	}

	return append(arr[split:], arr[:split]...)
}
