package main

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 4, 5),
			k:    0,
			want: newList(1, 2, 3, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			k:    3,
			want: newList(3, 4, 5, 1, 2),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			k:    6,
			want: newList(5, 1, 2, 3, 4),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			k:    1,
			want: newList(5, 1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := rotateRight(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
