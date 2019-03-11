package main

import (
	"fmt"
	"testing"
)

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

	for _, tt := range tests {
		z := tt.x + tt.y
		name := fmt.Sprintf("%d + %d == %d", tt.x, tt.y, z)
		t.Run(name, func(t *testing.T) {
			l1 := newListByDigit(tt.x)
			fmt.Printf("%13d | %s\n", l1.toInt(), l1)
			l2 := newListByDigit(tt.y)
			fmt.Printf("%13d | %s\n", l2.toInt(), l2)
			got := addTwoNumbers(l1, l2)
			actual := got.toInt()
			fmt.Printf("%5d ~ %5d | %s\n", z, actual, got)
			if actual != z {
				t.Errorf("addTwoNumbers() = %v, want %v", actual, z)
			}
		})
	}
}
