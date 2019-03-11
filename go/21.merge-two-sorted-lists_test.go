package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   newList(1),
			l2:   newList(2),
			want: newList(1, 2),
		},
		{
			l1:   newList(1),
			l2:   newList(2, 2),
			want: newList(1, 2, 2),
		},
		{
			l1:   newList(1, 2, 3),
			l2:   newList(2, 2, 4),
			want: newList(1, 2, 2, 2, 3, 4),
		},
		{
			l1:   newList(),
			l2:   newList(2, 2, 4),
			want: newList(2, 2, 4),
		},
		{
			l1:   newList(),
			l2:   newList(),
			want: newList(),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			fmt.Printf("l1: %s\n", tt.l1)
			fmt.Printf("l2: %s\n", tt.l2)
			got := mergeTwoLists(tt.l1, tt.l2)
			fmt.Printf("got: %s\n", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
