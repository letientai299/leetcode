package main

import "testing"

func Test_maxSumBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(4, 3, 8, 3, NA, NA, 7, 1, 4, NA, 6, NA, NA, NA, 5, 10), want: 13},
		{root: treeFromLevelOrder(1, 4, 3, 2, 4, 2, 5, NA, NA, NA, NA, NA, NA, 4, 6), want: 20},
		{root: treeFromLevelOrder(4, 3, NA, 1, 2), want: 2},
		{root: treeFromLevelOrder(-4, -2, -5), want: 0},
		{root: treeFromLevelOrder(2, 1, 3), want: 6},
		{root: treeFromLevelOrder(5, 4, 8, 3, NA, 6, 3), want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumBST(tt.root); got != tt.want {
				t.Errorf("maxSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
