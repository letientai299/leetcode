package main

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want []*ListNode
	}{
		{
			head: newList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			k:    3,
			want: []*ListNode{
				newList(1, 2, 3, 4),
				newList(5, 6, 7),
				newList(8, 9, 10),
			},
		},

		{
			head: newList(1, 2, 3),
			k:    2,
			want: []*ListNode{
				newList(1, 2),
				newList(3),
			},
		},

		{
			head: newList(1, 2, 3),
			k:    5,
			want: []*ListNode{
				newList(1),
				newList(2),
				newList(3),
				nil,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
