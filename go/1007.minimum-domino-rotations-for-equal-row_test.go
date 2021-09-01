package main

import "testing"

func Test_minDominoRotations(t *testing.T) {
	tests := []struct {
		name    string
		tops    []int
		bottoms []int
		want    int
	}{
		{
			tops:    []int{3, 5, 1, 2, 3},
			bottoms: []int{3, 6, 3, 3, 4},
			want:    -1,
		},

		{
			tops:    []int{2, 1, 2, 4, 2, 2},
			bottoms: []int{5, 2, 6, 2, 3, 2},
			want:    2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations(tt.tops, tt.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
