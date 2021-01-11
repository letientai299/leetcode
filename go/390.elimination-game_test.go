package main

import "testing"

func Test_lastRemaining(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 12, want: 6},
		{n: 10, want: 8},
		{n: 9, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.n); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
