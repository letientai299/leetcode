package main

import (
	"reflect"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []*TreeNode
	}{
		{
			n: 3, want: []*TreeNode{
				treeFromLevelOrder(0, 0, 0),
			},
		},


		{
			n: 1, want: []*TreeNode{
				treeFromLevelOrder(0),
			},
		},

		{n: 2, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
