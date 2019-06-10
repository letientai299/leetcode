package main

import (
	"fmt"
	"testing"
)

func Test_findTilt(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{treeFromLevelOrder(1, 2, 3, 4, NA, 5), 11},
		{treeFromLevelOrder(1, 2, 3), 1},
		{treeFromLevelOrder(1, 2), 2},
		{treeFromLevelOrder(1, 2, 3, 2), 3},
		{treeFromLevelOrder(1, 2, 3, 2, NA, NA, NA, 2), 9},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := findTilt(tt.root)
			if got != tt.want {
				t.Errorf("findTilt(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
