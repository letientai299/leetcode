package main

import "testing"

func Test_rankTeams(t *testing.T) {
	tests := []struct {
		votes []string
		want  string
	}{
		{
			votes: []string{"WXYZ", "XYZW"},
			want:  "XWYZ",
		},

		{
			votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			want:  "ACB",
		},

		{
			votes: []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"},
			want:  "ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := rankTeams(tt.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
