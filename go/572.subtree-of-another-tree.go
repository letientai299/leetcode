package main

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 *
 * https://leetcode.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (41.71%)
 * Total Accepted:    103.7K
 * Total Submissions: 248.2K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 *
 * Given two non-empty binary trees s and t, check whether tree t has exactly
 * the same structure and node values with a subtree of s. A subtree of s is a
 * tree consists of a node in s and all of this node's descendants. The tree s
 * could also be considered as a subtree of itself.
 *
 *
 * Example 1:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return true, because t has the same structure and node values with a subtree
 * of s.
 *
 *
 * Example 2:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 * ⁠   /
 * ⁠  0
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return false.
 *
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var isSameTree func(s *TreeNode, t *TreeNode) bool
	isSameTree = func(s *TreeNode, t *TreeNode) bool {
		if (s == nil) != (t == nil) {
			return false
		}

		if s == nil {
			return true
		}

		if s.Val != t.Val {
			return false
		}

		if (s.Left == nil) != (t.Left == nil) {
			return false
		}

		if (s.Right == nil) != (t.Right == nil) {
			return false
		}

		return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
	}

	if isSameTree(s, t) {
		return true
	}

	if s.Left != nil && isSubtree(s.Left, t) {
		return true
	}

	if s.Right != nil && isSubtree(s.Right, t) {
		return true
	}

	return false
}
