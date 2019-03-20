package main

import (
	"testing"
)

func Test_hasCycle_true(t *testing.T) {
	head := newList(1, 2, 3, 4)
	head.Next.Next.Next = head
	if yes := hasCycle(head); !yes {
		t.Errorf("hasCycle() = %v, want %v", yes, true)
	}
}

func Test_hasCycle_false(t *testing.T) {
	head := newList(1, 2, 3, 4)
	if yes := hasCycle(head); yes {
		t.Errorf("hasCycle() = %v, want %v", yes, false)
	}
}
