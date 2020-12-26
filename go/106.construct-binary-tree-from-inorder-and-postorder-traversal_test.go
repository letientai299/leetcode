package main

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		name      string
		postorder []int
		inorder   []int
		want      *TreeNode
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want:      treeFromLevelOrder(3, 9, 20, NA, NA, 15, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
