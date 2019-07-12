package main

import (
	"fmt"
	"testing"
)

func Test_findSecondMinimumValue(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(5, 8, 5),
			want: 8,
		},

		{
			root: treeFromLevelOrder(2, 1, 3),
			want: 2,
		},

		{
			root: treeFromLevelOrder(2, 2, 2),
			want: -1,
		},

		{
			root: treeFromLevelOrder(2, 2, 5, NA, NA, 5, 7),
			want: 5,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := findSecondMinimumValue(tt.root)
			if got != tt.want {
				t.Errorf("findSecondMinimumValue(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
