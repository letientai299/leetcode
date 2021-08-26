package main

// Complete Binary Tree Inserter
//
// Medium
//
// https://leetcode.com/problems/complete-binary-tree-inserter/
//
// A **complete binary tree** is a binary tree in which every level, except
// possibly the last, is completely filled, and all nodes are as far left as
// possible.
//
// Design an algorithm to insert a new node to a complete binary tree keeping it
// complete after the insertion.
//
// Implement the `CBTInserter` class:
//
// - `CBTInserter(TreeNode root)` Initializes the data structure with the `root`
// of the complete binary tree.
// - `int insert(int v)` Inserts a `TreeNode` into the tree with value `Node.val
// == val` so that the tree remains complete, and returns the value of the
// parent of the inserted `TreeNode`.
// - `TreeNode get_root()` Returns the root node of the tree.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/08/03/lc-treeinsert.jpg)
//
// ```
// Input
// ["CBTInserter", "insert", "insert", "get_root"]
// [[[1, 2]], [3], [4], []]
// Output
// [null, 1, 2, [1, 2, 3, 4]]
//
// Explanation
// CBTInserter cBTInserter = new CBTInserter([1, 2]);
// cBTInserter.insert(3);  // return 1
// cBTInserter.insert(4);  // return 2
// cBTInserter.get_root(); // return [1, 2, 3, 4]
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree will be in the range `[1, 1000]`.
// - `0 <= Node.val <= 5000`
// - `root` is a complete binary tree.
// - `0 <= val <= 5000`
// - At most `104` calls will be made to `insert` and `get_root`.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type CBTInserter struct {
	complete bool
	height   int
	left     *CBTInserter
	right    *CBTInserter
	root     *TreeNode
}

func Constructor919(root *TreeNode) CBTInserter {
	cb := CBTInserter{
		root:     root,
		height:   1,
		complete: true,
	}

	if root.Left != nil {
		child := Constructor919(root.Left)
		cb.left = &child
		if child.height >= cb.height {
			cb.height = child.height + 1
		}
	}

	if root.Right != nil {
		child := Constructor919(root.Right)
		cb.right = &child
		if child.height >= cb.height {
			cb.height = child.height + 1
		}
	}

	cb.complete = (cb.left.h() == cb.right.h()) &&
		(cb.left == nil || cb.left.complete) &&
		(cb.right == nil || cb.right.complete)

	return cb
}

func (cb *CBTInserter) h() int {
	if cb == nil {
		return 0
	}

	return cb.height
}

func (cb *CBTInserter) Insert(val int) int {
	defer func() {
		l, r := cb.left.h(), cb.right.h()
		if l < r {
			cb.height = r + 1
		} else {
			cb.height = l + 1
		}
		cb.complete = l == r &&
			(cb.left == nil || cb.left.complete) &&
			(cb.right == nil || cb.right.complete)
	}()

	if cb.left == nil || cb.right == nil {
		n := &TreeNode{Val: val}
		child := Constructor919(n)

		if cb.left == nil {
			cb.root.Left = n
			cb.left = &child
		} else {
			cb.root.Right = n
			cb.right = &child
		}

		return cb.root.Val
	}

	if !cb.left.complete || (cb.right.complete && cb.right.height == cb.left.height) {
		return cb.left.Insert(val)
	}

	return cb.right.Insert(val)
}

func (cb *CBTInserter) Get_root() *TreeNode {
	return cb.root
}
