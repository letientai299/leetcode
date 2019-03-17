package main

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		tree string
		want [][]int
	}{
		{
			tree: "",
			want: nil,
		},
		{
			tree: `{"val": 1}`,
			want: [][]int{{1}},
		},
		{
			tree: `
{
  "val": 3,
  "left": { "val": 9 },
  "right": {
    "val": 20,
    "left": { "val": 15 },
    "right": { "val": 7 }
  }
}
`,
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
		{
			tree: `{"val": 1}`,
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.tree, func(t *testing.T) {
			root := treeFromJson(tt.tree)
			if got := levelOrderBottom(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
