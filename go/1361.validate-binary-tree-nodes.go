package main

// Validate Binary Tree Nodes
//
// Medium
//
// https://leetcode.com/problems/validate-binary-tree-nodes/
//
// You have `n` binary tree nodes numbered from `0` to `n - 1` where node `i`
// has two children `leftChild[i]` and `rightChild[i]`, return `true` if and
// only if **all** the given nodes form **exactly one** valid binary tree.
//
// If node `i` has no left child then `leftChild[i]` will equal `-1`, similarly
// for the right child.
//
// Note that the nodes have no values and that we only use the node numbers in
// this problem.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2019/08/23/1503_ex1.png)
//
// ```
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
// Output: true
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2019/08/23/1503_ex2.png)
//
// ```
// Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
// Output: false
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2019/08/23/1503_ex3.png)
//
// ```
// Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]
// Output: false
//
// ```
//
// **Example 4:**
//
// ![](https://assets.leetcode.com/uploads/2019/08/23/1503_ex4.png)
//
// ```
// Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 104`
// - `leftChild.length == rightChild.length == n`
// - `-1 <= leftChild[i], rightChild[i] <= n - 1`
func validateBinaryTreeNodes(n int, lefts []int, rights []int) bool {
	type Node struct {
		val          int
		L, R, Parent *Node
	}

	notSame := func(a, b *Node) bool {
		return a != nil && b != nil && a != b
	}

	seen := make(map[int]*Node)

	for v, l := range lefts {
		r := rights[v]
		mid, ok := seen[v]
		if !ok {
			mid = &Node{val: v}
			seen[v] = mid
		}

		if (l == -1 && mid.L != nil) || (r == -1 && mid.R != nil) {
			return false
		}

		if l != -1 {
			left, ok := seen[l]
			if !ok {
				left = &Node{Parent: mid, val: l}
				seen[l] = left
				mid.L = left
			} else if notSame(mid.L, left) || notSame(left.Parent, mid) || mid.Parent == left {
				return false
			} else {
				left.Parent = mid
				mid.L = left
			}
		}

		if r != -1 {
			right, ok := seen[r]
			if !ok {
				right = &Node{Parent: mid, val: r}
				seen[r] = right
				mid.R = right
			} else if notSame(mid.R, right) || notSame(right.Parent, mid) || mid.Parent == right {
				return false
			} else {
				right.Parent = mid
				mid.R = right
			}
		}
	}

	// after all, there's only 1 Node without parent
	var root *Node
	for _, n := range seen {
		if n.Parent == nil {
			if root == nil {
				root = n
			} else {
				return false
			}
		}
	}

	if root == nil {
		return false
	}

	cnt := 0
	var walk func(n *Node)
	walk = func(node *Node) {
		if node == nil || cnt > n {
			return
		}

		cnt++
		walk(node.L)
		walk(node.R)
	}

	walk(root)
	return cnt == n
}

// TODO (tai): slow
