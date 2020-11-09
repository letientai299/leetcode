package main

import "testing"

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{arr: []int{0, 1, 2, 4, 2, 1}, want: true},
		{arr: []int{2, 0, 1}, want: false},
		{arr: []int{0, 3, 2, 1}, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := validMountainArray(tt.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
