package main

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{
			head: newList(1, 2, 3, 4, 5, 6, 7),
			m:    1,
			n:    7,
			want: newList(7, 6, 5, 4, 3, 2, 1),
		},

		{
			head: newList(1, 2, 3, 4, 5),
			m:    2,
			n:    4,
			want: newList(1, 4, 3, 2, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.head, tt.m, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
