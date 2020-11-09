package main

import "testing"

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{
		{a: []int{2, 1, 1, 1, 1, 2, 1}, want: false},
		{a: []int{2, 2, 2, 1}, want: true},
		{a: []int{1, 2, 2, 1}, want: false},
		{a: []int{1, 2, 1, 1}, want: false},
		{a: []int{1, 1, 1, 1}, want: true},
		{a: []int{1}, want: true},
		{a: []int{1, 2, 2, 3}, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isMonotonic(tt.a); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
