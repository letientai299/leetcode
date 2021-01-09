package main

// medium
// 1372. Longest ZigZag Path in a Binary Tree
func longestZigZag(root *TreeNode) int {
	var zig func(n *TreeNode) (left, right int)

	best := 1

	zig = func(n *TreeNode) (left, right int) {
		if n == nil {
			return 0, 0
		}

		ll, lr := zig(n.Left)
		rl, rr := zig(n.Right)
		if best < ll {
			best = ll
		}
		if best < rr {
			best = rr
		}
		if best < 1+lr {
			best = 1 + lr
		}
		if best < 1+rl {
			best = 1 + rl
		}
		return 1 + lr, 1 + rl
	}

	zig(root)
	return best - 1
}
