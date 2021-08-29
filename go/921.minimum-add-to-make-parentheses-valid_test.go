package main

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: ")(", want: 2},
		{s: "()(", want: 1},
		{s: ")()", want: 1},
		{s: "))(", want: 3},
		{s: "()", want: 0},
		{s: "", want: 0},
		{s: "(", want: 1},
		{s: ")", want: 1},
		{s: ")()()()()((", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minAddToMakeValid(tt.s); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
