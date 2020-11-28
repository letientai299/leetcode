package main

import "testing"

func Test_solveEquation(t *testing.T) {
	tests := []struct {
		equation string
		want     string
	}{
		{equation: "0x=0", want: "Infinite solutions"},
		{equation: "x+5-3+x=6+x-2", want: "x=2"},
		{equation: "x=2+x", want: "No solution"},
		{equation: "x=x", want: "Infinite solutions"},
		{equation: "x+2=4", want: "x=2"},
	}
	for _, tt := range tests {
		t.Run(tt.equation, func(t *testing.T) {
			if got := solveEquation(tt.equation); got != tt.want {
				t.Errorf("solveEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
