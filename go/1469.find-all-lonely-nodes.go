package main

func getLonelyNodes(root *TreeNode) []int {
	var res []int

	var check func(n *TreeNode)

	check = func(n *TreeNode) {
		if n == nil {
			return
		}

		if (n.Left == nil) != (n.Right == nil) {
			if n.Left != nil {
				res = append(res, n.Left.Val)
			} else {
				res = append(res, n.Right.Val)
			}
		}

		check(n.Left)
		check(n.Right)
	}

	check(root)
	return res
}
