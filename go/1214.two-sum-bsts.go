package main

// Two Sum BSTs
//
// # Medium
//
// https://leetcode.com/problems/two-sum-bsts
func twoSumBSTs(a *TreeNode, b *TreeNode, t int) bool {
	m := make(map[int]bool)

	var collect func(n *TreeNode)
	collect = func(n *TreeNode) {
		if n != nil {
			m[n.Val] = true
			collect(n.Left)
			collect(n.Right)
		}
	}

	collect(a)

	var walk func(*TreeNode) bool
	walk = func(n *TreeNode) bool {
		if n == nil {
			return false
		}

		v := t - n.Val
		return m[v] || walk(n.Left) || walk(n.Right)
	}

	return walk(b)
}
