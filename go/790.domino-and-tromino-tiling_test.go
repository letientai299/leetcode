package main

import "testing"

func Test_numTilings(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{                  // 0
		{n: 1, want: 1},  // 0*2 + 1
		{n: 2, want: 2},  // 1*2 + 0
		{n: 3, want: 5},  // 2*2 + 1
		{n: 4, want: 11}, // 10 + 1
		{n: 5, want: 24}, // 22 + 2
		{n: 6, want: 53}, // 48 + 5
		{n: 7, want: 117},// 106 + 11
		{n: 8, want: 258},// 234 + 24
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilings(tt.n); got != tt.want {
				t.Errorf("numTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
