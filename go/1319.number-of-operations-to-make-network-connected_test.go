package main

import "testing"

func Test_makeConnected(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        int
	}{
		{
			n:    12,
			want: 4,
			connections: [][]int{
				{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2},
			},
		},

		{
			n:           4,
			want:        1,
			connections: [][]int{{0, 1}, {0, 2}, {1, 2}},
		},

		{
			n:           5,
			connections: [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}},
			want:        0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.n, tt.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
