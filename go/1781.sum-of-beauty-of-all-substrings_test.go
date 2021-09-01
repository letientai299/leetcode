package main

import "testing"

func Test_beautySum(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "aabcb", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := beautySum(tt.s); got != tt.want {t.Errorf("beautySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
