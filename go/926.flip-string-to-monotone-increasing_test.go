package main

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "00110", want: 1},
		{s: "010110", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
