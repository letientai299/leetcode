package main

import "math"

// Minimum Cost Tree From Leaf Values
//
// Medium
//
// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
//
// Given an array `arr` of positive integers, consider all binary trees such
// that:
//
// - Each node has either `0` or `2` children;
// - The values of `arr` correspond to the values of each **leaf** in an
// in-order traversal of the tree.
// - The value of each non-leaf node is equal to the product of the largest leaf
// value in its left and right subtree, respectively.
//
// Among all possible binary trees considered, return _the smallest possible sum
// of the values of each non-leaf node_. It is guaranteed this sum fits into a
// **32-bit** integer.
//
// A node is a **leaf** if and only if it has zero children.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/10/tree1.jpg)
//
// ```
// Input: arr = [6,2,4]
// Output: 32
// Explanation: There are two possible trees shown.
// The first has a non-leaf node sum 36, and the second has non-leaf node sum
// 32.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/10/tree2.jpg)
//
// ```
// Input: arr = [4,11]
// Output: 44
//
// ```
//
// **Constraints:**
//
// - `2 <= arr.length <= 40`
// - `1 <= arr[i] <= 15`
// - It is guaranteed that the answer fits into a **32-bit** signed integer
// (i.e., it is less than 231).
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	var walk func(l, r int) (sum int, root int)
	minSums := make([][]int, 0, n)
	leaves := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		minSums = append(minSums, make([]int, n+1))
		leaves = append(leaves, make([]int, n+1))
	}

	walk = func(l, r int) (int, int) {
		if leaf := leaves[l][r]; leaf != 0 {
			return minSums[l][r], leaf
		}

		if r-l == 1 {
			leaves[l][r] = arr[l]
			return 0, arr[l]
		}

		if r-l == 2 {
			s := arr[l] * arr[l+1]
			v := arr[l]
			if arr[l+1] > v {
				v = arr[l+1]
			}

			leaves[l][r] = v
			minSums[l][r] = s
			return s, v
		}

		minSum := math.MaxInt32
		leaf := 0
		for i := l + 1; i < r; i++ {
			sumLeft, largestLeft := walk(l, i)
			sumRight, largestRight := walk(i, r)
			root := largestLeft * largestRight
			if sumLeft+sumRight+root < minSum {
				minSum = sumLeft + sumRight + root
				leaf = largestLeft
				if leaf < largestRight {
					leaf = largestRight
				}
			}
		}

		leaves[l][r] = leaf
		minSums[l][r] = minSum
		return minSum, leaf
	}

	min, _ := walk(0, len(arr))
	return min
}
