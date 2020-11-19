package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 4),
			n:    1,
			want: newList(1, 2, 3),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    1,
			want: newList(1, 2, 3, 4),
		},

		{
			head: newList(1, 2),
			n:    1,
			want: newList(1),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    6,
			want: newList(1, 2, 3, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    5,
			want: newList(2, 3, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    4,
			want: newList(1, 3, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    3,
			want: newList(1, 2, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			n:    2,
			want: newList(1, 2, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
