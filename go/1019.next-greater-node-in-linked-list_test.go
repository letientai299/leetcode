package main

import (
	"reflect"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want []int
	}{
		{
			head: newList(1, 7, 9, 2, 5, 1),
			want: []int{7, 9, 0, 5, 0, 0},
		},

		{
			head: newList(2, 2, 5),
			want: []int{5, 5, 0},
		},

		{
			// s: 9 7 6 7
			// r: 0 0 7 0
			head: newList(9, 7, 6, 7, 6, 9),
			want: []int{0, 9, 7, 9, 9, 0},
		},

		{
			head: newList(1, 7, 5, 1, 9, 2, 5, 1),
			want: []int{7, 9, 9, 9, 0, 5, 0, 0},
		},

		{
			head: newList(2, 7, 4, 3, 5),
			want: []int{7, 0, 5, 5, 0},
		},

		{
			head: newList(2, 1, 5, 2),
			want: []int{5, 5, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
