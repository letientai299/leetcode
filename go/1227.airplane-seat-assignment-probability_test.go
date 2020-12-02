package main

import "testing"

func Test_nthPersonGetsNthSeat(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want float64
	}{
		{n: 1, want: 1},
		{n: 2, want: 0.5},
		{n: 3, want: 0.5},
		{n: 4, want: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPersonGetsNthSeat(tt.n); got != tt.want {
				t.Errorf("nthPersonGetsNthSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
