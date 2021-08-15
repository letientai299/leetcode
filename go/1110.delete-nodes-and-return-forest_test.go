package main

import (
	"reflect"
	"testing"
)

func Test_delNodes(t *testing.T) {
	tests := []struct {
		name string
		root     *TreeNode
		toDelete []int
		want []*TreeNode
	}{
		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7),
			toDelete: []int{3, 5},
			want: []*TreeNode{
				treeFromLevelOrder(1,2,NA, 4),
				treeFromLevelOrder(6),
				treeFromLevelOrder(7),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNodes(tt.root, tt.toDelete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
