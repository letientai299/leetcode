package main

import "testing"

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name  string
		cards []int
		k     int
		want  int
	}{
		{cards: []int{11, 49, 100, 20, 86, 29, 72}, k: 4, want: 232},
		{cards: []int{100, 40, 17, 9, 73, 75}, k: 3, want: 248},
		{cards: []int{1, 2, 3}, k: 1, want: 3},
		{cards: []int{1, 2, 3, 4, 5, 6, 1}, k: 3, want: 12},
		{cards: []int{1, 6, 2, 3}, k: 2, want: 7},
		{cards: []int{1, 1, 4, 1, 1, 1}, k: 3, want: 6},
		{cards: []int{1, 6, 2, 3}, k: 4, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.cards, tt.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
