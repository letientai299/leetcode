package main

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 4, 5),
			k:    3,
			want: newList(3, 2, 1, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			k:    6,
			want: newList(1, 2, 3, 4, 5),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			k:    2,
			want: newList(2, 1, 4, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
