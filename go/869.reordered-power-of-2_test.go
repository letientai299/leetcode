package main

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{n: 821, want: true},
		{n: 9, want: false},
		{n: 8, want: true},
		{n: 4, want: true},
		{n: 2, want: true},
		{n: 10, want: false},
		{n: 1, want: true},
		{n: 46, want: true},
		{n: 47, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
