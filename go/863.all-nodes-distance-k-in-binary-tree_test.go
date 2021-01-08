package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distanceK(t *testing.T) {
	type args struct {
		root   *TreeNode
		target *TreeNode
		K      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root:   treeFromLevelOrder(0, 1, NA, NA, 2, NA, 3, NA, 4),
				target: treeFromLevelOrder(3),
				K:      0,
			},
			want: []int{3},
		},

		{
			args: args{
				root:   treeFromLevelOrder(0, 1, NA, 3, 2),
				target: treeFromLevelOrder(2),
				K:      1,
			},
			want: []int{1},
		},

		{
			args: args{
				root:   treeFromLevelOrder(3, 5, 1, 6, 2, 0, 8, NA, NA, 7, 4, NA, NA, NA, NA, 10, 9, NA, 11),
				target: treeFromLevelOrder(5),
				K:      3,
			},
			want: []int{10, 9, 11, 0, 8},
		},

		{
			args: args{
				root:   treeFromLevelOrder(3, 5, 1, 6, 2, 0, 8, NA, NA, 7, 4),
				target: treeFromLevelOrder(5),
				K:      2,
			},
			want: []int{7, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, tt.args.target, tt.args.K); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
				tt.args.root.print()
			}
		})
	}
}
