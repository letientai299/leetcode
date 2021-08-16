package main

import "testing"

func Test_minimumLengthEncoding(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			words: []string{"ti", "me", "bell", "time"},
			want:  13,
		},

		{
			words: []string{"t"},
			want:  2,
		},

		{
			words: []string{"time", "ti", "me", "bell"},
			want:  13,
		},

		{
			words: []string{"time", "me", "bell"},
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding(tt.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
