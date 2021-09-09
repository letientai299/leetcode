package main

import "testing"

func Test_countHomogenous(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abbcccaa", want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := countHomogenous(tt.s); got != tt.want {
				t.Errorf("countHomogenous() = %v, want %v", got, tt.want)
			}
		})
	}
}
