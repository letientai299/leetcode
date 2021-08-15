package main

import "testing"

func Test_alphabetBoardPath(t *testing.T) {
	tests := []struct {
		target string
		want   string
	}{
		{
			target: "leet", want: "DDR!UURRR!!DDD!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			if got := alphabetBoardPath(tt.target); got != tt.want {
				t.Errorf("alphabetBoardPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
