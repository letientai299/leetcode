package main

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	ms := MsConstructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)

	fmt.Println(ms.mins)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.mins)
	fmt.Println(ms.GetMin())
}
