package main

import (
	"fmt"
	"testing"
)

func Test_smallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		k    int
		want int
	}{
		{k: 0, want: -1},
		{k: 1, want: 1},
		{k: 2, want: -1},
		{k: 3, want: 3},
		{k: 4, want: -1},
		{k: 5, want: -1},
		{k: 6, want: -1},
		{k: 7, want: 6},
		{k: 8, want: -1},
		{k: 9, want: 9},
		{k: 10, want: -1},
		{k: 11, want: 2},
		{k: 12, want: -1},
		{k: 13, want: 6},
		{k: 137, want: 8},
		{k: 5882353, want: 16},
		{k: 1234567, want: 34020},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.k), func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.k); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
