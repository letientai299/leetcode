package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []string
	}{
		{
			root: treeFromJson(`
{
  "val": 3,
  "left": { "val": 9 },
  "right": {
    "val": 20,
    "left": { "val": 15 },
    "right": { "val": 7 }
  }
}
`),
			want: []string{"3->9", "3->20->15", "3->20->7"},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			if got := binaryTreePaths(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
