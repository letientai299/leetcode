package main

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "2-3", want: -1},
		{s: "2+3", want: 5},
		{s: "2+3*4/5 + 1", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
