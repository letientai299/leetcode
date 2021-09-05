package main

import "testing"

func Test_findKthBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		// 0111001
		{n: 3, k: 3, want: 1},

		{n: 3, k: 1, want: 0},
		{n: 3, k: 2, want: 1},
		{n: 3, k: 4, want: 1},
		{n: 3, k: 5, want: 0},
		{n: 3, k: 6, want: 0},
		{n: 3, k: 7, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := byte(tt.want + '0')
			if got := findKthBit(tt.n, tt.k); got != want {
				t.Errorf("findKthBit() = %v, want %v", string(got), string(want))
			}
		})
	}
}
