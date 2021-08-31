package main

import "testing"

func Test_maxAncestorDiff(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(8, 3, 10, 1, 6, NA, 14, NA, NA, 4, 7, 13), want: 7},
		{root: treeFromLevelOrder(1, NA, 2, NA, 0, 3), want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxAncestorDiff(tt.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
