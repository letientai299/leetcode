package main

import (
	"reflect"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(2, 1, 3),
			want: treeFromLevelOrder(1, NA, 2, NA, 3),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := increasingBST(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
