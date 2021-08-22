package main

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists []*ListNode
		want  *ListNode
	}{
		{
			lists: []*ListNode{
				nil,
				nil,
			},
			want: nil,
		},

		{
			lists: []*ListNode{
				newList(1, 4, 5),
				newList(1, 3, 4),
				newList(2, 6),
			},
			want: newList(1, 1, 2, 3, 4, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
