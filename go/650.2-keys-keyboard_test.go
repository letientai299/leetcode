package main

import "testing"

func Test_minSteps(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 6, want: 5},
		{n: 120, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
