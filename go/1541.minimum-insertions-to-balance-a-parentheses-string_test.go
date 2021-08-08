package main

import "testing"

func Test_minInsertions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "))())(", want: 3},
		{s: "())", want: 0},
		{s: "((((((", want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
