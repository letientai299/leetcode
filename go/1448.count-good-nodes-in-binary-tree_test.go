package main

import "testing"

func Test_goodNodes(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(3, 1, 4, 3, NA, 1, 5), want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
