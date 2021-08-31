package main

import "testing"

func Test_smallestFromLeaf(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{root: treeFromLevelOrder(0, 1, 2, 3, 4, 3, 4), want: "dba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestFromLeaf(tt.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
