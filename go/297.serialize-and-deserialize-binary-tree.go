package main

import (
	"strconv"
	"strings"
)

// Serialize and Deserialize Binary Tree
//
// Hard
//
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
//
// Serialization is the process of converting a data structure or object into a
// sequence of bits so that it can be stored in a file or memory buffer, or
// transmitted across a network connection link to be reconstructed later in the
// same or another computer environment.
//
// Design an algorithm to serialize and deserialize a binary tree. There is no
// restriction on how your serialization/deserialization algorithm should work.
// You just need to ensure that a binary tree can be serialized to a string and
// this string can be deserialized to the original tree structure.
//
// **Clarification:** The input/output format is the same as [how LeetCode
// serializes a binary tree](/faq/#binary-tree). You do not necessarily need to
// follow this format, so please be creative and come up with different
// approaches yourself.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg)
//
// ```
// Input: root = [1,2,3,null,null,4,5]
// Output: [1,2,3,null,null,4,5]
//
// ```
//
// **Example 2:**
//
// ```
// Input: root = []
// Output: []
//
// ```
//
// **Example 3:**
//
// ```
// Input: root = [1]
// Output: [1]
//
// ```
//
// **Example 4:**
//
// ```
// Input: root = [1,2]
// Output: [1,2]
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[0, 104]`.
// - `-1000 <= Node.val <= 1000`

// Apparently my generic solution for 449 (same problem but with BST instead
// of generic tree) doesn't work. So here my second try.

type Codec297 struct{}

func Constructor297() Codec { return Codec{} }

func (cx *Codec297) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var sb strings.Builder
	// level order
	q := []*TreeNode{root}
	var next []*TreeNode
	nils := 0

	for len(q) != 0 {
		for _, n := range q {
			if n == nil {
				nils++
				continue
			}

			switch nils {
			case 0:
				break
			case 1:
				sb.WriteString("*,")
			default:
				s := strconv.Itoa(nils)
				sb.WriteString("*")
				sb.WriteString(s)
				sb.WriteString(",")
			}

			s := strconv.Itoa(n.Val)
			sb.WriteString(s)
			sb.WriteByte(',')
			next = append(next, n.Left)
			next = append(next, n.Right)
			nils = 0
		}

		q, next = next, q
		next = next[:0]
	}

	s := sb.String()
	n := len(s)
	for n--; n >= 0 && s[n] == ',' || s[n] == '*'; n-- {
	}

	return s[:n+1]
}

func (cx *Codec297) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	i := 0
	v := 0
	sig := 1
	for ; i < len(data) && data[i] != ','; i++ {
		if data[i] == '-' {
			sig = -1
			continue
		}
		v = 10*v + int(data[i]-'0')
	}

	root := &TreeNode{Val: sig * v}
	q := []*TreeNode{root}
	var next []*TreeNode

	nils := 0

	nextNode := func() *TreeNode {
		if i >= len(data) {
			return nil
		}

		if nils > 0 {
			nils--
			return nil
		}

		isNil := false
		v := 0
		sig := 1
		for i++; i < len(data) && data[i] != ','; i++ {
			if data[i] == '*' {
				isNil = true
				continue
			}

			if data[i] == '-' {
				sig = -1
				continue
			}

			v = 10*v + int(data[i]-'0')
		}

		if isNil {
			nils = v - 1
			return nil
		}

		return &TreeNode{Val: sig * v}
	}

	for len(q) != 0 {
		for j := 0; j < len(q) && i < len(data); j++ {
			n := q[j]
			if sub := nextNode(); sub != nil {
				n.Left = sub
				next = append(next, sub)
			}
			if sub := nextNode(); sub != nil {
				n.Right = sub
				next = append(next, sub)
			}
		}

		next, q = q, next
		next = next[:0]
	}
	return root
}
