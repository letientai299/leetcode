package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(4, 2, 7, 1, 3),
			val:  2,
			want: treeFromLevelOrder(2, 1, 3),
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.root, tt.val,
		)
		t.Run(testName, func(t *testing.T) {
			if got := searchBST(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST(%v, %v) = %v, want %v", tt.root, tt.val, got, tt.want)
			}
		})
	}
}
