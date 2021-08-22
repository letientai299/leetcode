package main

import "testing"

func Test_minDeletionSize(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{strs: []string{"abx","agz","bgc","bfc"}, want: 1},
		{strs: []string{"zyx", "wvu", "tsr"}, want: 3},
		{strs: []string{"xc", "yb", "za"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
