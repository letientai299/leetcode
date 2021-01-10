package main

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 4, 5),
			k:    2,
			want: newList(1, 4, 3, 2, 5),
		},

		{
			head: newList(7, 9, 6, 6, 7, 8, 3, 0, 9, 5),
			k:    5,
			want: newList(7, 9, 6, 6, 8, 7, 3, 0, 9, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5, 6),
			k:    2,
			want: newList(1, 5, 3, 4, 2, 6),
		},

		{
			head: newList(1, 2, 3, 4, 5, 6),
			k:    4,
			want: newList(1, 2, 4, 3, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
