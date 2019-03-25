package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		{newList(1), true},
		{newList(1, 2, 2, 1), true},
		{newList(1, 2, 4, 2, 1), true},
		{newList(1, 2, 4, 3, 1), false},
		{newList(1, 2, 4, 2, 2), false},
		{newList(1, 2, 1), true},
		{newList(1, 2), false},
		{newList(), true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.head,
		)
		t.Run(testName, func(t *testing.T) {
			got := isPalindrome(tt.head)
			if got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.head, got, tt.want)
			}
		})
	}
}
