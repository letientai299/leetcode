package main

import "testing"

func Test_kthFactor(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{n: 4, k: 4, want: -1},
		{n: 124, k: 4, want: 31},
		{n: 124, k: 1, want: 1},
		{n: 124, k: 30, want: -1},
		{n: 124, k: 3, want: 4},
		{n: 12, k: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthFactor(tt.n, tt.k); got != tt.want {
				t.Errorf("kthFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
