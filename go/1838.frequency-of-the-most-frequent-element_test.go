package main

import "testing"

func Test_maxFrequency(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 4}, k: 5, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
