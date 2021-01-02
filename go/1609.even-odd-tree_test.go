package main

import "testing"

func Test_isEvenOddTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			root: treeFromLevelOrder(2, 12, 8, 5, 9, NA, NA, 18, 16),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
				tt.root.print()
			}
		})
	}
}
