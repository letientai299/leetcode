package main

import (
	"fmt"
	"testing"
)

func Test_findTarget(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want bool
	}{
		{
			root: treeFromLevelOrder(2, 1, 3),
			k:    4,
			want: true,
		},

		{
			root: treeFromLevelOrder(5, 3, 6, 2, 4, NA, 7),
			k:    9,
			want: true,
		},

		{
			root: treeFromLevelOrder(5, 3, 6, 2, 4, NA, 7),
			k:    28,
			want: false,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.root, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := findTarget(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("findTarget(%v, %v) = %v, want %v", tt.root, tt.k, got, tt.want)
			}
		})
	}
}
