package main

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	tests := []struct {
		name   string
		tokens []int
		power  int
		want   int
	}{
		{
			power:  200,
			tokens: []int{100, 200, 300, 400},
			want:   2,
		},

		{
			power:  51,
			tokens: []int{26},
			want:   1,
		},

		{
			power:  50,
			tokens: []int{100},
			want:   0,
		},

		{
			power:  150,
			tokens: []int{100, 200},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.tokens, tt.power); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
