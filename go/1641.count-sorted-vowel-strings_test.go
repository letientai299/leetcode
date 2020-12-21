package main

import "testing"

func Test_countVowelStrings(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 33, want: 66045},
		{n: 2, want: 15},
		{n: 1, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelStrings(tt.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
