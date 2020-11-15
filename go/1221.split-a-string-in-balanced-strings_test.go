package main

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "RLRRLLRLRL", want: 4},
		{s: "RLLR", want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := balancedStringSplit(tt.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
