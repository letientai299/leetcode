package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 0),
			x:    -1,
			want: newList(1, 2, 3, 0),
		},

		{
			head: newList(1, 2, 3, 0),
			x:    5,
			want: newList(1, 2, 3, 0),
		},

		{
			head: newList(1, 2, 3, 0),
			x:    2,
			want: newList(1, 0, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := partition(tt.head, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
