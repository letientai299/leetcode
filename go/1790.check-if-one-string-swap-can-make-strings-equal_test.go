package main

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{s1: "bnak", s2: "kanb", want: false},
		{s1: "bank", s2: "kanb", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.s1, tt.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
