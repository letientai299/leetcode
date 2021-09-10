package main

import (
	"fmt"
	"testing"
)

func TestAllOne(t *testing.T) {
	a := Constructor432()
	a.Inc("hello")
	a.Inc("hello")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
	a.Inc("leet")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())

	a.Dec("hello")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())

	a.Dec("leet")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())

	a.Dec("hello")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
}
