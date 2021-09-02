package main

import "testing"

func Test_minNumberOfFrogs(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "crocakcroraoakk", want: 2},
		{s: "croak", want: 1},
		{s: "craok", want: -1},
		{s: "croa", want: -1},
		{s: "croakcroak", want: 1},
		{s: "croackroak", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minNumberOfFrogs(tt.s); got != tt.want {
				t.Errorf("minNumberOfFrogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
