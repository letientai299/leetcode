package main

import "testing"

func Test_numTeams(t *testing.T) {
	tests := []struct {
		name   string
		rating []int
		want   int
	}{
		{rating: []int{2, 5, 3, 4, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTeams(tt.rating); got != tt.want {
				t.Errorf("numTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
