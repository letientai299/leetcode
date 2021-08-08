package main

import "testing"

func Test_tupleSameProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{2, 3, 4, 6, 8, 12}, want: 40},
		{nums: []int{2, 3, 4, 6}, want: 8},
		{nums: []int{1, 2, 4, 5, 10}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
