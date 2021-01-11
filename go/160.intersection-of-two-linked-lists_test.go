package main

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	tests := []struct {
		a *ListNode
		b *ListNode
	}{
		{
			a: newList(1),
			b: newList(5, 6),
		},

		{
			a: newList(1, 2, 3, 4),
			b: newList(5, 6),
		},

		{
			a: newList(1, 2, 3, 4, 5),
			b: newList(5, 6),
		},

		{
			a: newList(1),
			b: newList(5),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.b
			for want.Next != nil {
				want = want.Next
			}

			middle := tt.a
			fast := tt.a.Next
			for fast != nil {
				middle = middle.Next
				fast = fast.Next
				if fast == nil {
					break
				}

				fast = fast.Next
			}

			want.Next = middle
			want = middle
			if got := getIntersectionNode(tt.a, tt.b); got != want {
				t.Errorf("getIntersectionNode() = %v, want %v", got, want)
			}
		})
	}
}

func Test_getIntersectionNode_no_intersection(t *testing.T) {
	tests := []struct {
		a *ListNode
		b *ListNode
	}{
		{
			a: newList(1),
			b: newList(5, 6),
		},

		{
			a: newList(1, 2, 3, 4),
			b: newList(5, 6),
		},

		{
			a: newList(1, 2, 3, 4, 5),
			b: newList(5, 6),
		},

		{
			a: newList(),
			b: newList(5),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			fmt.Printf("a (before): %v\n", tt.a)
			fmt.Printf("b (before): %v\n", tt.b)
			if got := getIntersectionNode(tt.a, tt.b); got != nil {
				t.Errorf("getIntersectionNode() = %v, want %v", got, nil)
			}
			fmt.Printf("a (after): %v\n", tt.a)
			fmt.Printf("b (after): %v\n", tt.b)
		})
	}
}
