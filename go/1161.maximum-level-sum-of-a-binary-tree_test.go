package main

import "testing"

func Test_maxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(-100, -200, -300, -20, -5, -10, NA),
			want: 3,
		},

		{
			root: treeFromLevelOrder(989, NA, 10250, 98693, -89388, NA, NA, NA, -32127),
			want: 2,
		},

		{
			root: treeFromLevelOrder(1, 7, 0, 7, -8),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
