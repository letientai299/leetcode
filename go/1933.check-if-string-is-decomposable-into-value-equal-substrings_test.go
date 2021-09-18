package main

import "testing"

func Test_isDecomposable(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "001", want: false},
		{s: "66666666666677722", want: true},
		{s: "000111000", want: false},
		{s: "00011111222", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isDecomposable(tt.s); got != tt.want {
				t.Errorf("isDecomposable() = %v, want %v", got, tt.want)
			}
		})
	}
}
