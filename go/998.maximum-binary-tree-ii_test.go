package main

import (
	"reflect"
	"testing"
)

func Test_insertIntoMaxTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(5, 2, 4, NA, 1),
			val:  3,
			want: treeFromLevelOrder(5, 2, 4, NA, 1, NA, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoMaxTree(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoMaxTree() = %v, want %v", got, tt.want)
				tt.want.print()
				t.Log("actual: ")
				got.print()
			}
		})
	}
}
