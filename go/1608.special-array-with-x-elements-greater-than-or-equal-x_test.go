package main

import "testing"

func Test_specialArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{2, 3}, want: 2},
		{nums: []int{0, 0, 3, 4, 4}, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := specialArray(tt.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
