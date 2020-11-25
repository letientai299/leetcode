package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: newList(1, 2, 3),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
