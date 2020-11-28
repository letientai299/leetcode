package main

import "testing"

func Test_numRabbits(t *testing.T) {
	tests := []struct {
		name    string
		answers []int
		want    int
	}{
		{answers: []int{3, 1}, want: 6},
		{answers: []int{1}, want: 2},
		{answers: []int{0}, want: 1},
		{answers: []int{}, want: 0},
		{answers: []int{1, 2, 2, 1, 1, 1}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
