package main

import "testing"

func Test_tictactoe(t *testing.T) {
	tests := []struct {
		moves [][]int
		want  string
	}{
		{
			moves: [][]int{{0, 0}, {1, 1}},
			want:  "Pending",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tictactoe(tt.moves); got != tt.want {
				t.Errorf("tictactoe() = %v, want %v", got, tt.want)
			}
		})
	}
}
