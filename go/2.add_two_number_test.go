package main

import (
	"fmt"
	"strings"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	buf := strings.Builder{}
	for {
		buf.WriteString(fmt.Sprintf("%d", n.Val))
		if n.Next != nil {
			n = n.Next
			buf.WriteString(" --> ")
		} else {
			break
		}
	}

	return buf.String()
}

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{
			x: 12345,
			y: 300,
		},
		{
			x: 0,
			y: 300,
		},
		{
			x: 12345,
			y: 98765,
		},
		{
			x: 243,
			y: 564,
		},
	}
	toList := func(x int) *ListNode {
		root := new(ListNode)
		n := root
		for {
			n.Val = x % 10
			x = x / 10
			if x > 0 {
				n.Next = new(ListNode)
				n = n.Next
			} else {
				break
			}
		}
		return root
	}

	toInt := func(n *ListNode) int {
		v := 0
		p := 1
		for n != nil {
			v = v + p*n.Val
			p = p * 10
			n = n.Next
		}
		return v
	}

	for _, tt := range tests {
		z := tt.x + tt.y
		name := fmt.Sprintf("%d + %d == %d", tt.x, tt.y, z)
		t.Run(name, func(t *testing.T) {
			l1 := toList(tt.x)
			fmt.Printf("%13d | %s\n", toInt(l1), l1)
			l2 := toList(tt.y)
			fmt.Printf("%13d | %s\n", toInt(l2), l2)
			got := addTwoNumbers(l1, l2)
			actual := toInt(got)
			fmt.Printf("%5d ~ %5d | %s\n", z, actual, got)
			if actual != z {
				t.Errorf("addTwoNumbers() = %v, want %v", actual, z)
			}
		})
	}
}
