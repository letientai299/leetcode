package main

import (
	"fmt"
	"testing"
)

func Test_longestUnivaluePath(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(1, 1, 2, 1), want: 2},
		{root: treeFromLevelOrder(1, 4, 5, 4, 4, 5), want: 2},
		{root: treeFromLevelOrder(5, 4, 5, 1, 1, 5), want: 2},
		{root: treeFromLevelOrder(), want: 0},
		{root: treeFromLevelOrder(1), want: 0},
		{root: treeFromLevelOrder(1, 1), want: 1},
		{root: treeFromLevelOrder(1, 2, 3, 1), want: 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := longestUnivaluePath(tt.root)
			if got != tt.want {
				t.Errorf("longestUnivaluePath(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
