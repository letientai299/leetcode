package main

import "testing"

func Test_kthGrammar(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{n: 30, k: 434991989, want: 0},
		{n: 3, k: 4, want: 0},
		{n: 3, k: 3, want: 1},
		{n: 3, k: 2, want: 1},
		{n: 3, k: 1, want: 0},
		{n: 2, k: 2, want: 1},
		{n: 2, k: 1, want: 0},
		{n: 1, k: 1, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.n, tt.k); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
