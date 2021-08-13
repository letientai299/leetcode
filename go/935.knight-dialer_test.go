package main

import "testing"

func Test_knightDialer(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 3131, want: 136006598},
		{n: 1, want: 10},
		{n: 2, want: 20},
		{n: 3, want: 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightDialer(tt.n); got != tt.want {
				t.Errorf("knightDialer() = %v, want %v", got, tt.want)
			}
		})
	}
}
