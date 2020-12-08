package main

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{head: newList(2, 1, 3, 0), want: newList(0, 1, 2, 3)},
		{head: newList(5, 4, 2, 1), want: newList(1, 2, 4, 5)},
		{head: newList(2, 1, 3), want: newList(1, 2, 3)},
		{head: newList(2, 1), want: newList(1, 2)},
		{head: newList(), want: newList()},
		{head: newList(1), want: newList(1)},
		{head: newList(1, 2, 3, 0), want: newList(0, 1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
