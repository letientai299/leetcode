package main

import "testing"

func Test_concatenatedBinary(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 12, want: 505379714},
		{n: 3, want: 27},
		{n: 2, want: 6},
		{n: 1, want: 1},
		{n: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenatedBinary(tt.n); got != tt.want {
				t.Errorf("concatenatedBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
