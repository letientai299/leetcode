package main

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(1, NA, 0, 0, 1),
			want: treeFromLevelOrder(1, NA, 0, NA, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
				tt.want.print()
				t.Log("actual: ")
				got.print()
			}
		})
	}
}
