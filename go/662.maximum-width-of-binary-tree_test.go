package main

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(1, 1, 1, 1, NA, NA, 1, 1, NA, NA, 1),
			want: 8,
		},

		{
			root: treeFromLevelOrder(1, 3, NA, 5, 3),
			want: 2,
		},

		{
			root: treeFromLevelOrder(1, 3, 2, 5, 3, NA, 9),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
				tt.root.print()
			}
		})
	}
}
