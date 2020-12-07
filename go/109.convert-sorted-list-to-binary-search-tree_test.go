package main

import (
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
	}{
		{head: newList(1, 2, 3, 4, 5, 6)},
		{head: newList(1, 2, 3, 4, 5)},
		{head: newList(1, 2, 3)},
		{head: newList(1, 2)},
		{head: newList(1)},
		{head: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedListToBST(tt.head)
			if !isValidBST(got) || !isBalanced(got) {
				t.Errorf("sortedListToBST() = %v", got)
			}
		})
	}
}
