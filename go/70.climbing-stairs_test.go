package main

import (
	"strconv"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		// f(4) = 1/2  * (f(3) * f(1) + f(2) * f(2) + f(1) * f(3)) ==  5
		// f(4) = 1 + 1 + 1 + 1
		// 		= 2 + 1 + 1
		// 		= 1 + 2 + 1
		// 		= 1 + 1 + 2
		// 		= 2 + 2
		{4, 5},
		// f(5) = 1 + 1 + 1 + 1 + 1				(1)
		//      = 2 + 1 + 1 + 1 				(4)
		//      = 2 + 2 + 1 					(3)
		{5, 8},
		// f(6) = 1 + 1 + 1 + 1 + 1 + 1  		(1)
		// 		= 2 + 1 + 1 + 1 + 1				(5)
		// 		= 2 + 2 + 1 + 1					(6)
		// 		= 2 + 2 + 2						(1)
		{6, 13},
		{44, 1134903170},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
