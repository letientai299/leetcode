package main

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		v    int
		d    int
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(4, 2, NA, 3, 1),
			v:    1,
			d:    3,
			want: treeFromLevelOrder(4, 2, NA, 1, 1, 3, NA, NA, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.root, tt.v, tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
