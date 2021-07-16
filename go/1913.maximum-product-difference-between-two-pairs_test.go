package main

import "testing"

func Test_maxProductDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference(tt.nums); got != tt.want {
				t.Errorf("maxProductDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
