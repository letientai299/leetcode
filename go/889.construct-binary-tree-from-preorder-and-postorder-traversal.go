package main

/*
 * @lc app=leetcode id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (62.96%)
 * Total Accepted:    44.7K
 * Total Submissions: 66.5K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * Return any binary tree that matches the given preorder and postorder
 * traversals.
 *
 * Values in the traversals pre and post are distinct positive integers.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
 * Output: [1,2,3,4,5,6,7]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= pre.length == post.length <= 30
 * pre[] and post[] are both permutations of 1, 2, ..., pre.length.
 * It is guaranteed an answer exists. If there exists multiple answers, you can
 * return any of them.
 *
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}

	if len(pre) == 1 {
		return root
	}

	i := 0
	for ; i < len(post); i++ {
		if post[i] == pre[1] {
			break
		}
	}

	root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
	root.Right = constructFromPrePost(pre[i+2:], post[i+1:len(post)-1])
	return root
}
