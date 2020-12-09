package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7, 8),
			p:    &TreeNode{Val: 9},
			q:    &TreeNode{Val: 5},
			want: nil,
		},

		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7, 8),
			p:    &TreeNode{Val: 8},
			q:    &TreeNode{Val: 5},
			want: &TreeNode{Val: 2},
		},

		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 5},
			want: &TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if tt.want == nil {
				assert.Nil(t, got)
				return
			}

			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
