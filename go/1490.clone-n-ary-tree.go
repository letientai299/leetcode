package main

// Clone N-ary Tree
//
// # Medium
//
// https://leetcode.com/problems/clone-n-ary-tree
func cloneTree(root *Node) *Node {
	cloned := make(map[*Node]*Node)

	var clone func(n *Node) *Node
	clone = func(n *Node) *Node {
		if n == nil {
			return n
		}

		if v, ok := cloned[n]; ok {
			return v
		}

		v := &Node{Val: n.Val}
		cloned[n] = v
		for _, c := range n.Children {
			v.Children = append(v.Children, clone(c))
		}

		return v
	}

	return clone(root)
}
