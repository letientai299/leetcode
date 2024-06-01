package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitBST(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			args: args{
				root:   treeFromLevelOrder(10, 5, 20, 3, 9, 15, 25, NA, NA, 8, NA, NA, NA, NA, NA, 6, NA, NA, 7),
				target: 6,
			},
			want: []*TreeNode{
				treeFromLevelOrder(5, 3, 6),
				treeFromLevelOrder(10, 9, 20, 8, NA, 15, 25, 7),
			},
		},
		{
			args: args{
				root:   treeFromLevelOrder(4, 2, 6, 1, 3, 5, 7),
				target: 2,
			},

			want: []*TreeNode{
				treeFromLevelOrder(2, 1),
				treeFromLevelOrder(4, 3, 6, NA, NA, 5, 7),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, splitBST(tt.args.root, tt.args.target), "splitBST(%v, %v)", tt.args.root, tt.args.target)
		})
	}
}
