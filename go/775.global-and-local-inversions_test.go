package main

import "testing"

func Test_isIdealPermutation(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want bool
	}{
		{a: []int{2, 1, 3, 0, 5, 4}, want: false},
		{a: []int{0, 1, 3, 2, 5, 4}, want: true},
		{a: []int{1, 0, 3, 2, 5, 4}, want: true},
		{a: []int{1, 0, 3, 2}, want: true},
		{a: []int{1, 0, 2, 3}, want: true},
		{a: []int{2, 0, 1, 3}, want: false},
		{a: []int{2, 1, 0, 3}, want: false},
		{a: []int{0, 2, 1, 3}, want: true},
		{a: []int{0, 1, 2, 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIdealPermutation(tt.a); got != tt.want {
				t.Errorf("isIdealPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
