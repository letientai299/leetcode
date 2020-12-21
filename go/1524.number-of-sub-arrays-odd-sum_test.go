package main

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, want: 16},
		{arr: []int{100, 100, 99, 99}, want: 4},
		{arr: []int{2, 4, 6}, want: 0},
		{arr: []int{1, 3, 5}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.arr); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
