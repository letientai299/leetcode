package main

import "testing"

func Test_squareIsWhite(t *testing.T) {
	tests := []struct {
		coordinates string
		want        bool
	}{
		{coordinates: "a1"},
		{coordinates: "a3"},
		{coordinates: "e1"},
		{coordinates: "g1"},
		{coordinates: "a2", want: true},
		{coordinates: "a4", want: true},
		{coordinates: "e4", want: true},
		{coordinates: "g2", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.coordinates, func(t *testing.T) {
			if got := squareIsWhite(tt.coordinates); got != tt.want {
				t.Errorf("squareIsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}
