package main

import "testing"

func Test_countTriples(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 4, want: 0},
		{n: 5, want: 2},
		{n: 10, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples(tt.n); got != tt.want {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}
