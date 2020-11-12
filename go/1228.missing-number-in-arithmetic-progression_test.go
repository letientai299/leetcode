package main

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{arr: []int{1, 3, 4}, want: 2},
		{arr: []int{4, 2, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := missingNumber(tt.arr); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
