package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSumBSTs(t *testing.T) {
	type args struct {
		a *TreeNode
		b *TreeNode
		t int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				a: treeFromLevelOrder(0, -10, 10),
				b: treeFromLevelOrder(5, 1, 7, 0, 2),
				t: 17,
			},
			want: true,
		},

		{
			args: args{
				a: treeFromLevelOrder(0, -10, 10),
				b: treeFromLevelOrder(5, 1, 7, 0, 2),
				t: 18,
			},
			want: false,
		},

		{
			args: args{
				a: treeFromLevelOrder(2, 1, 4),
				b: treeFromLevelOrder(1, 0, 3),
				t: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, twoSumBSTs(tt.args.a, tt.args.b,
				tt.args.t), "twoSumBSTs(%v, %v, %v)", tt.args.a, tt.args.b,
				tt.args.t)
		})
	}
}
