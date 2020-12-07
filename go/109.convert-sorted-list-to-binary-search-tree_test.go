package main

import (
	"reflect"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *TreeNode
	}{
		{
			head: newList(1, 2, 3, 4), want: treeFromLevelOrder(3, 2, 4, 1),
		},

		{
			head: newList(1, 2, 3, 4, 5), want: treeFromLevelOrder(3, 2, 5, 1, NA, 4),
		},

		{
			head: newList(1, 2), want: treeFromLevelOrder(2, 1),
		},

		{
			head: newList(1, 2, 3), want: treeFromLevelOrder(2, 1, 3),
		},
		{
			head: newList(1), want: &TreeNode{Val: 1},
		},
		{
			head: nil, want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
