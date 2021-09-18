package main

import (
	"fmt"
	"testing"
)

func Test_consecutiveNumbersSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 5, want: 2},
		{n: 9, want: 3},
		{n: 15, want: 4},
		{n: 341231232, want: 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.n); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
