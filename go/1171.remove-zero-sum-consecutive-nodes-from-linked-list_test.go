package main

import (
	"reflect"
	"testing"
)

func Test_removeZeroSumSublists(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			head: newList(1, 3, 2, -3, -2, 5, 5, -5, 1),
			want: newList(1, 5, 1),
		},

		{
			head: newList(1, 2, 3, -3, -2),
			want: newList(1),
		},

		{
			head: newList(1, 0, 2, 3, -3, -1, 0, -1, 3),
			want: newList(1, 3),
		},

		{
			head: newList(1, 2, -3, 3, 1),
			want: newList(3, 1),
		},

		{
			head: newList(1, 2, 3, -3, 4),
			want: newList(1, 2, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroSumSublists(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
	}
}
