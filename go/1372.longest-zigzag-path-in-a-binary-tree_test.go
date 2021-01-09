package main

import "testing"

func Test_longestZigZag(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(1, 1, 1, NA, 1, NA, NA, 1, 1, NA, 1),
			want: 4,
		},

		{
			root: treeFromLevelOrder(1, NA, 1, 1, 1, NA, NA, 1, 1, NA, 1, NA, NA, NA, 1, NA, 1),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestZigZag(tt.root); got != tt.want {
				t.Errorf("longestZigZag() = %v, want %v", got, tt.want)
			}
		})
	}
}
