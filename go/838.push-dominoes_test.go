package main

import "testing"

func Test_pushDominoes(t *testing.T) {
	tests := []struct {
		dominoes string
		want     string
	}{
		{dominoes: "L.....RR.RL.....L.R.", want: "L.....RRRRLLLLLLL.RR"},
		{dominoes: "RR.L", want: "RR.L"},
		{dominoes: ".L.R...LR..L..", want: "LL.RR.LLRRLL.."},
	}
	for _, tt := range tests {
		t.Run(tt.dominoes, func(t *testing.T) {
			if got := pushDominoes(tt.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
