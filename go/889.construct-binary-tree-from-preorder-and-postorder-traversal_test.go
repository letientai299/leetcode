package main

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	tests := []struct {
		want *TreeNode
		name string
	}{
		{
			want: treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pre := preorderTraversal(tt.want)
			post := postorderTraversal(tt.want)

			if got := constructFromPrePost(pre, post); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
				t.Log("pre: ", pre)
				t.Log("post: ", post)
				tt.want.print(func(s string) { t.Log(s) })
			}
		})
	}
}
