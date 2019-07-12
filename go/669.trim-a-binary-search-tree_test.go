package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		L    int
		R    int
		want *TreeNode
	}{

		{
			root: treeFromLevelOrder(3, 1, 4, NA, 2),
			L:    1,
			R:    2,
			want: treeFromLevelOrder(1, NA, 2),
		},

		{
			root: treeFromLevelOrder(3, 0, 4, NA, 2, NA, NA, 1),
			L:    1,
			R:    3,
			want: treeFromLevelOrder(3, 2, NA, 1),
		},

		{
			root: treeFromLevelOrder(1, 0, 2),
			L:    1,
			R:    2,
			want: treeFromLevelOrder(1, NA, 2),
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v",
			tt.root, tt.L, tt.R,
		)
		t.Run(testName, func(t *testing.T) {
			root := treeFromLevelOrder(tt.root.levelOrder()...)
			if got := trimBST(root, tt.L, tt.R); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST(%v, %v, %v) = %v, want %v", tt.root, tt.L, tt.R, got, tt.want)
			}
		})
	}
}
