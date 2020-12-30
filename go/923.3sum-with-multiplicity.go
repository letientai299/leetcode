package main

/*
 * @lc app=leetcode id=923 lang=golang
 *
 * [923] 3Sum With Multiplicity
 *
 * https://leetcode.com/problems/3sum-with-multiplicity/description/
 *
 * algorithms
 * Medium (34.98%)
 * Total Accepted:    21.6K
 * Total Submissions: 60.2K
 * Testcase Example:  '[1,1,2,2,3,3,4,4,5,5]\n8'
 *
 * Given an integer array A, and an integer target, return the number of tuples
 * i, j, k  such that i < j < k and A[i] + A[j] + A[k] == target.
 *
 * As the answer can be very large, return it modulo 10^9 + 7.
 *
 *
 * Example 1:
 *
 *
 * Input: A = [1,1,2,2,3,3,4,4,5,5], target = 8
 * Output: 20
 * Explanation:
 * Enumerating by the values (A[i], A[j], A[k]):
 * (1, 2, 5) occurs 8 times;
 * (1, 3, 4) occurs 8 times;
 * (2, 2, 4) occurs 2 times;
 * (2, 3, 3) occurs 2 times.
 *
 *
 * Example 2:
 *
 *
 * Input: A = [1,1,2,2,2,2], target = 5
 * Output: 12
 * Explanation:
 * A[i] = 1, A[j] = A[k] = 2 occurs 12 times:
 * We choose one 1 from [1,1] in 2 ways,
 * and two 2s from [2,2,2,2] in 6 ways.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= A.length <= 3000
 * 0 <= A[i] <= 100
 * 0 <= target <= 300
 *
 *
 */
func threeSumMulti(nums []int, target int) int {
	m := make(map[int]int)
	var arr []int
	res := 0
	for _, v := range nums {
		if m[v] == 0 {
			arr = append(arr, v)
		}
		m[v]++
	}

	if target%3 == 0 && m[target/3] >= 3 {
		x := m[target/3]
		res += x * (x - 1) * (x - 2) / 6
		res %= 1e9 + 7
	}

	all := threeSumGeneral(arr, target)
	for _, x := range all {
		res += m[x[0]] * m[x[1]] * m[x[2]]
		res %= 1e9 + 7
	}

	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})

	for _, v := range arr {
		x1 := target - 2*v
		if _, ok := m1[x1]; ok {
			//fmt.Println("2v", v, x1)
			res += m[x1] * m[v] * (m[v] - 1) / 2
			res %= 1e9 + 7
		} else {
			m1[v] = struct{}{}
		}

		x2 := target - v
		if _, ok := m2[x2]; ok {
			//fmt.Println("1v", v, x2/2)
			res += m[v] * m[x2/2] * (m[x2/2] - 1) / 2
			res %= 1e9 + 7
		} else {
			m2[2*v] = struct{}{}
		}
	}

	return res
}
