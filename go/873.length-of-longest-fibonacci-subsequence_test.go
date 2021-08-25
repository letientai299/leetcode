package main

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenLongestFibSubseq(tt.arr); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
