package main

import "testing"

func Test_numTilePossibilities(t *testing.T) {
	tests := []struct {
		tiles string
		want  int
	}{
		{tiles: "AAABBC", want: 188},
		{tiles: "AV", want: 4},
		{tiles: "A", want: 1},
		{tiles: "AA", want: 2},
		{tiles: "AAB", want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.tiles, func(t *testing.T) {
			if got := numTilePossibilities(tt.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
