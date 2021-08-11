package main

import (
	"reflect"
	"testing"
)

func Test_removeLeafNodes(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		target int
		want   *TreeNode
	}{
		{
			root:   treeFromLevelOrder(1, 2, 1),
			target: 2,
			want:   treeFromLevelOrder(1, NA, 1),
		},

		{
			root:   treeFromLevelOrder(1, 1, 1),
			target: 1,
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.root, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
