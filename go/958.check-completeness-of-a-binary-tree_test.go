package main

import "testing"

func Test_isCompleteTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{root: treeFromLevelOrder(1, 2, 3, 4), want: true},
		{root: treeFromLevelOrder(1, 2), want: true},
		{root: treeFromLevelOrder(1, 2, NA, 4), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompleteTree(tt.root); got != tt.want {
				t.Errorf("isCompleteTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
