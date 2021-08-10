package main

import "testing"

func Test_maximalNetworkRank(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		roads [][]int
		want  int
	}{
		{
			n: 4,
			roads: [][]int{
				{0, 1}, {0, 3}, {1, 2}, {1, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalNetworkRank(tt.n, tt.roads); got != tt.want {
				t.Errorf("maximalNetworkRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
