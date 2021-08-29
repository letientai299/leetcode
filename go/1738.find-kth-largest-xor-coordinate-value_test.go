package main

import "testing"

func Test_kthLargestValue(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		k    int
		want int
	}{
		{
			mat: [][]int{
				{10, 9, 5}, {2, 0, 4}, {1, 0, 9}, {3, 4, 8},
			},
			k:    10,
			want: 1,
		},
		{
			mat: [][]int{{5, 2}, {1, 6}},
			// XORed
			// 5 7
			// 4 0
			k:    1,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.mat, tt.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
