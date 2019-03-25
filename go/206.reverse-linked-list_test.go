package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: newList(1, 2, 3),
			want: newList(3, 2, 1),
		},
		{
			head: newList(1),
			want: newList(1),
		},
		{
			head: nil,
			want: nil,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.head,
		)
		t.Run(testName, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList(%v) = %v, want %v", tt.head, got, tt.want)
			}
		})
	}
}
