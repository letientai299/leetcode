package main

import "testing"

func Test_superPow(t *testing.T) {
	tests := []struct {
		a    int
		b    []int
		want int
	}{
		{a: 2147483647, b: []int{2, 0, 0}, want: 1198},
		{a: 2, b: []int{2, 0, 0}, want: 1215},
		{a: 2, b: []int{2, 0}, want: 368},
		{a: 2, b: []int{1, 0}, want: 1024},
		{a: 0, b: []int{1, 0}, want: 0},
		{a: 1, b: []int{1, 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := superPow(tt.a, tt.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
