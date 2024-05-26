package main

// Recover Binary Search Tree
//
// Medium
//
// https://leetcode.com/problems/recover-binary-search-tree
//
// You are given the `root` of a binary search tree (BST), where the values of
// **exactly** two nodes of the tree were swapped by mistake. _Recover the tree
// without changing its structure_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/10/28/recover1.jpg)
//
// ```
// Input: root = [1,3,null,null,2]
// Output: [3,1,null,null,2]
// Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3
// makes the BST valid.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/10/28/recover2.jpg)
//
// ```
// Input: root = [3,1,4,null,null,2]
// Output: [2,1,4,null,null,3]
// Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2
// and 3 makes the BST valid.
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[2, 1000]`.
// - `-231 <= Node.val <= 231 - 1`
//
// **Follow up:** A solution using `O(n)` space is pretty straight-forward.
// Could you devise a constant `O(1)` space solution?
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// improve from v2 by storing only the previous node in the inorder list.
// O(n) time and O(1) space
func recoverTree(root *TreeNode) {
	var left *TreeNode  // track the first (inorder) faulty node
	var right *TreeNode // track the last faulty node.
	var pre *TreeNode   // track the previous node in the inorder

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		if pre != nil && node.Val < pre.Val {
			if left == nil {
				left = pre
			}
			right = node
		}
		pre = node
		inorder(node.Right)
	}

	inorder(root)
	left.Val, right.Val = right.Val, left.Val
}

// improve from v1 by tracking left and right faulty nodes during traversal.
// O(n) time and O(n) space
func recoverTree_v2(root *TreeNode) {
	var arr []*TreeNode
	var left *TreeNode  // track the first (inorder) faulty node
	var right *TreeNode // track the last faulty node.

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		if len(arr) != 0 {
			last := arr[len(arr)-1]
			if node.Val < last.Val {
				if left == nil {
					left = last
				}
				right = node
			}
		}
		arr = append(arr, node)
		inorder(node.Right)
	}

	inorder(root)
	left.Val, right.Val = right.Val, left.Val
}

// inorder traversal, collect all values, fix it with O(n) time.
// O(4n) time and O(n) space.
func recoverTree_v1(root *TreeNode) {
	var arr []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			arr = append(arr, node.Val)
			inorder(node.Right)
		}
	}

	inorder(root)

	var left int
	for left = 0; left < len(arr)-1 && arr[left] < arr[left+1]; left++ {
	}
	var right int
	for right = len(arr) - 1; right > 0 && arr[right-1] < arr[right]; right-- {
	}
	arr[left], arr[right] = arr[right], arr[left]

	var fix func(node *TreeNode)
	i := 0
	fix = func(node *TreeNode) {
		if node != nil {
			fix(node.Left)
			node.Val = arr[i]
			i++
			fix(node.Right)
		}
	}

	fix(root)
}
