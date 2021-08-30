package main

import "testing"

func Test_kConcatenationMaxSum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{arr: []int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000}, k: 100000, want: 999999937},
		{arr: []int{-5, -2, 0, 0, 3, 9, -2, -5, 4}, k: 5, want: 20},
		{arr: []int{2, -1, 2, -2}, k: 4, want: 6},
		{arr: []int{-1, -2}, k: 4, want: 0},
		{arr: []int{2, -1, -2}, k: 4, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kConcatenationMaxSum(tt.arr, tt.k); got != tt.want {
				t.Errorf("kConcatenationMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
