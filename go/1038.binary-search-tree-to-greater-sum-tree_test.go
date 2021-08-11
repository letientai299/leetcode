package main

import (
	"reflect"
	"testing"
)

func Test_bstToGst(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(4, 1, 6, 0, 2, 5, 7, NA, NA, NA, 3, NA, NA, NA, 8),
			want: treeFromLevelOrder(30, 36, 21, 36, 35, 26, 15, NA, NA, NA, 33, NA, NA, NA, 8),
		},

		{
			root: treeFromLevelOrder(2, 1, 3),
			want: treeFromLevelOrder(5, 6, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstToGst(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}
