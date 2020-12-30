package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=449 lang=golang
 *
 * [449] Serialize and Deserialize BST
 *
 * https://leetcode.com/problems/serialize-and-deserialize-bst/description/
 *
 * algorithms
 * Medium (49.55%)
 * Total Accepted:    139K
 * Total Submissions: 258.9K
 * Testcase Example:  '[2,1,3]'
 *
 * Serialization is converting a data structure or object into a sequence of
 * bits so that it can be stored in a file or memory buffer, or transmitted
 * across a network connection link to be reconstructed later in the same or
 * another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary search tree. There
 * is no restriction on how your serialization/deserialization algorithm should
 * work. You need to ensure that a binary search tree can be serialized to a
 * string, and this string can be deserialized to the original tree structure.
 *
 * The encoded string should be as compact as possible.
 *
 *
 * Example 1:
 * Input: root = [2,1,3]
 * Output: [2,1,3]
 * Example 2:
 * Input: root = []
 * Output: []
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^4].
 * 0 <= Node.val <= 10^4
 * The input tree is guaranteed to be a binary search tree.
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

type Codec struct{}

func Constructor449() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var sb strings.Builder
	row := []*TreeNode{root}
	for len(row) != 0 {
		next := make([]*TreeNode, 0, 2*len(row))
		for _, n := range row {
			if n == nil {
				sb.WriteByte('#')
				sb.WriteByte(',')
				continue
			}
			s := strconv.Itoa(n.Val)
			sb.WriteString(s)
			sb.WriteByte(',')
			next = append(next, n.Left)
			next = append(next, n.Right)
		}

		row = next
	}

	s := sb.String()
	i := len(s) - 1
	for i >= 0 && (s[i] == ',' || s[i] == '#') {
		i--
	}
	return s[:i+1]
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	v := 0
	i := 0
	for i < len(data) && data[i] != ',' {
		v = 10*v + int(data[i]-'0')
		i++
	}

	root := &TreeNode{
		Val: v,
	}

	if i == len(data) {
		return root
	}

	var row []*TreeNode
	left := true
	node := root
	v = 0
	for i++; i < len(data); i++ {
		x := data[i]
		switch x {
		case ',': // finish value
			t := &TreeNode{Val: v}
			row = append(row, t)
			if left {
				node.Left = t
			} else {
				node.Right = t
				node = row[0]
				row = row[1:]
			}
			left = !left
			v = 0
		case '#': // nil node
			if !left {
				node = row[0]
				row = row[1:]
			}
			left = !left
			v = 0
			i++
		default:
			v = 10*v + int(x-'0')
		}
	}

	t := &TreeNode{Val: v}
	if left {
		node.Left = t
	} else {
		node.Right = t
	}

	return root
}
