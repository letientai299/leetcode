package main

import "testing"

func Test_isBipartite(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{
			graph: [][]int{
				{4}, {}, {4}, {4}, {0, 2, 3},
			},
			want: true,
		},

		{
			graph: [][]int{
				{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8},
			},
			want: false,
		},

		{
			graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			want:  false,
		},

		{
			graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
