package main

import "testing"

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(2, -1),
			want: 2,
		},

		{
			root: treeFromLevelOrder(-10, 9, 20, NA, NA, 15, 7),
			want: 42,
		},

		{
			root: treeFromLevelOrder(1, 2, 3),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
