package main

import "testing"

func Test_maximumSwap(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1993, want: 9913},
		{n: 9976, want: 9976},
		{n: 2736, want: 7236},
		{n: 99234, want: 99432},
		{n: 23456, want: 63452},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumSwap(tt.n); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
