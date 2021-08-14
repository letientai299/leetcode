package main

import "testing"

func Test_numSteps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "1101", want: 6},
		{s: "1000", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numSteps(tt.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
