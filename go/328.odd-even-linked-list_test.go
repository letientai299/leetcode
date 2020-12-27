package main

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{head: newList(2, 1, 3, 5, 6, 4, 7), want: newList(2, 3, 6, 7, 1, 5, 4)},
		{head: newList(2, 1, 3, 5), want: newList(2, 3, 1, 5)},
		{head: newList(2, 1), want: newList(2, 1)},
		{head: newList(2), want: newList(2)},
		{head: newList(1), want: newList(1)},
		{head: nil, want: nil},

		{
			head:
			newList(1, 2, 3, 4, 5),
			want: newList(1, 3, 5, 2, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
