package main

import (
	"fmt"
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{treeFromLevelOrder(1, NA, 3, 2), 1},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := getMinimumDifference(tt.root)
			if got != tt.want {
				t.Errorf("getMinimumDifference(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
