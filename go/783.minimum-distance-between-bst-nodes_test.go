package main

import (
	"fmt"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(90, 69, NA, 49, 89, NA, 52, NA, NA, NA, NA),
			want: 1,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := minDiffInBST(tt.root)
			if got != tt.want {
				t.Errorf("minDiffInBST(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
