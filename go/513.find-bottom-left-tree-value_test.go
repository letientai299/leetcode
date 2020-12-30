package main

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(1, 2, 3, 4, NA, 5, 6, NA, NA, 7),
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
