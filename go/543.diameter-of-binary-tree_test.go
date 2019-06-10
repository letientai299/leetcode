package main

import (
	"fmt"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{treeFromLevelOrder(1, 2, 3, 4, 5, NA), 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := diameterOfBinaryTree(tt.root)
			if got != tt.want {
				t.Errorf("diameterOfBinaryTree(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
