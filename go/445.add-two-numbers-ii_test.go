package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   newList(7, 2, 4, 3),
			l2:   newList(5, 6, 4),
			want: newList(7, 8, 0, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
