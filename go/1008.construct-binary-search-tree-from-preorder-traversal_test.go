package main

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     *TreeNode
	}{
		{
			preorder: []int{8, 5, 1, 7, 10, 12},
			want:     treeFromLevelOrder(8, 5, 10, 1, 7, NA, 12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
