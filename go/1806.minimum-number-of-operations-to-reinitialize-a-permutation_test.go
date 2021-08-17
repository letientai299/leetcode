package main

import (
	"testing"
)

func Test_reinitializePermutation(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 2, want: 1},
		{n: 4, want: 2},
		{n: 6, want: 4},
		{n: 100, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reinitializePermutation(tt.n); got != tt.want {
				t.Errorf("reinitializePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
