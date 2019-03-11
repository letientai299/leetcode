package main

import (
	"fmt"
	"strings"
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

func newList(arr ...int) *ListNode {
	if	len(arr) == 0 {
		return nil
	}

	r := new(ListNode)
	n := r
	for i, x := range arr {
		if i == len(arr)-1 {
			break
		}
		n.Val = x
		n.Next = new(ListNode)
		n = n.Next
	}

	n.Val = arr[len(arr)-1]
	return r
}

func newListByDigit(x int) *ListNode {
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

func (n *ListNode) toInt() int {
	v := 0
	p := 1
	for n != nil {
		v = v + p*n.Val
		p = p * 10
		n = n.Next
	}
	return v
}
