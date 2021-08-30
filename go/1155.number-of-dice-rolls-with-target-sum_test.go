package main

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	tests := []struct {
		name   string
		d      int
		f      int
		target int
		want   int
	}{
		{d: 30, f: 30, target: 500, want: 222616187},
		{d: 2, f: 6, target: 7, want: 6},
		{d: 1, f: 6, target: 3, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.d, tt.f, tt.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
