package main

import "testing"

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{root: treeFromLevelOrder(3, 1, 4, NA, 2), k: 4, want: 4},
		{root: treeFromLevelOrder(3, 1, 4, NA, 2), k: 3, want: 3},
		{root: treeFromLevelOrder(3, 1, 4, NA, 2), k: 2, want: 2},
		{root: treeFromLevelOrder(3, 1, 4, NA, 2), k: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
