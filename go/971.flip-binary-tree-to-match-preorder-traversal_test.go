package main

import (
	"reflect"
	"testing"
)

func Test_flipMatchVoyage(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		voyage []int
		want   []int
	}{
		{
			root:   treeFromLevelOrder(1, 2, 3),
			voyage: []int{1, 3, 2},
			want:   []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipMatchVoyage(tt.root, tt.voyage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipMatchVoyage() = %v, want %v", got, tt.want)
			}
		})
	}
}
