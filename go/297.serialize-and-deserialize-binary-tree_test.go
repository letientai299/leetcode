package main

import (
	"reflect"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{
			root: treeFromLevelOrder(
				4, -7, -3, NA, NA, -9, -3, 9, -7, -4, NA, 6, NA, -6, -6, NA, NA, 0, 6, 5, NA, 9, NA, NA, -1, -4, NA, NA, NA, -2,
			),
		},
		{root: treeFromLevelOrder(1, 2, 3, NA, 4)},
		{root: treeFromLevelOrder(1, 2, 3, NA, NA, NA, 4)},
		{root: treeFromLevelOrder(1, 2, 3, NA, NA, NA, 4, 5)},
		{root: treeFromLevelOrder(1, 2, -3, NA, NA, NA, -4, -5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cx := &Codec{}
			data := cx.serialize(tt.root)
			got := cx.deserialize(data)
			if !reflect.DeepEqual(got, tt.root) {
				t.Errorf("deserialize() = %v, want %v", got, tt.root)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			root: treeFromLevelOrder(
				4, -7, -3, NA, NA, -9, -3, 9, -7, -4, NA, 6, NA, -6, -6, NA, NA, 0, 6, 5, NA, 9, NA, NA, -1, -4, NA, NA, NA, -2,
			),
			want: "4,-7,-3,*2,-9,-3,9,-7,-4,*,6,*,-6,-6,*2,0,6,5,*,9,*2,-1,-4,*3,-2",
		},
		{root: treeFromLevelOrder(1, -2, 3, NA, -4), want: "1,-2,3,*,-4"},
		{root: treeFromLevelOrder(1, 2, 3, NA, 4), want: "1,2,3,*,4"},
		{root: treeFromLevelOrder(1, 2, 3, NA, NA, NA, 4), want: "1,2,3,*3,4"},
		{root: treeFromLevelOrder(1, 2, 3, NA, NA, NA, 4, 5), want: "1,2,3,*3,4,5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cx := &Codec{}
			if got := cx.serialize(tt.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
