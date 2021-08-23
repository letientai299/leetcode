package main

import "testing"

func Test_pseudoPalindromicPaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(2, 1, 1, 1, 3, NA, NA, NA, NA, NA, 1), want: 1},
		{root: treeFromLevelOrder(2, 3, 1, 3, 1, NA, 1), want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
