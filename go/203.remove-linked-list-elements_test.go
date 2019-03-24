package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	tests := []struct {
		head *ListNode
		val  int
		want *ListNode
	}{
		{
			head: newListByDigit(3456788),
			val:  8,
			want: newListByDigit(34567),
		},

		{
			head: newListByDigit(12345678),
			val:  1,
			want: newListByDigit(2345678),
		},

		{
			head: nil,
			val:  1,
			want: nil,
		},

		{
			head: newListByDigit(1),
			val:  1,
			want: nil,
		},

		{
			head: newListByDigit(1),
			val:  0,
			want: newListByDigit(1),
		},

		{
			head: newListByDigit(111),
			val:  1,
			want: nil,
		},

		{
			head: newListByDigit(123456),
			val:  6,
			want: newListByDigit(12345),
		},

		{
			head: newListByDigit(121156),
			val:  1,
			want: newListByDigit(256),
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.head, tt.val,
		)
		t.Run(testName, func(t *testing.T) {
			if got := removeElements(tt.head, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements(%v, %v) = %v, want %v", tt.head, tt.val, got, tt.want)
			}
		})
	}
}
