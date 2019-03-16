package main

import (
	"strconv"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{9, 3},
		{10, 3},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.x), func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
