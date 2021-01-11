package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_printTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]string
	}{
		{
			root: treeFromLevelOrder(3, 1, 5, 0, 2, 4, 6, NA, NA, NA, 3),
			want: [][]string{
				{"", "", "", "", "", "", "", "3", "", "", "", "", "", "", ""},
				{"", "", "", "1", "", "", "", "", "", "", "", "5", "", "", ""},
				{"", "0", "", "", "", "2", "", "", "", "4", "", "", "", "6", ""},
				{"", "", "", "", "", "", "3", "", "", "", "", "", "", "", ""},
			},
		},

		{
			root: treeFromLevelOrder(1, 2, 5, 3, NA, NA, NA, 4),
			want: [][]string{
				{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
				{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
				{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
				{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			},
		},

		{
			root: treeFromLevelOrder(1),
			want: [][]string{{"1"}},
		},

		{
			root: nil,
			want: nil,
		},

		{
			root: treeFromLevelOrder(1, 2, 3, NA, 4),
			want: [][]string{
				{"", "", "", "1", "", "", ""},
				{"", "2", "", "", "", "3", ""},
				{"", "", "4", "", "", "", ""},
			},
		},

		{
			root: treeFromLevelOrder(1, 2),
			want: [][]string{
				{"", "1", ""},
				{"2", "", ""},
			},
		},

		{
			root: treeFromLevelOrder(1, NA, 2),
			want: [][]string{
				{"", "1", ""},
				{"", "", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Error("printTree() failed, want")
				for _, row := range tt.want {
					t.Log(strings.Join(row, "_"))
				}

				t.Log("actual: ")
				for _, row := range got {
					t.Log(strings.Join(row, "_"))
				}
			}
		})
	}
}
