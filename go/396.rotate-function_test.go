package main

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{a: []int{4, 3, 2, 6}, want: 26},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxRotateFunction(tt.a); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
