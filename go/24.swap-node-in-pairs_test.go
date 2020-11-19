package main

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: nil,
			want: nil,
		},

		{
			head: newList(1),
			want: newList(1),
		},

		{
			head: newList(1, 2, 3),
			want: newList(2, 1, 3),
		},

		{
			head: newList(1, 2, 3, 4),
			want: newList(2, 1, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := swapPairs(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
