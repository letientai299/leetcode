package main

import "testing"

func Test_numberOfMatches(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 7, want: 6},
		{n: 5, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatches(tt.n); got != tt.want {
				t.Errorf("numberOfMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
