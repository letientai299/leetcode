package main

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	tests := []struct {
		name       string
		properties [][]int
		want       int
	}{
		{
			properties: [][]int{
				{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10},
			},
			want: 2,
		},

		{
			properties: [][]int{{5, 5}, {6, 3}, {3, 6}},
			want:       0,
		},

		{
			properties: [][]int{{1, 5}, {10, 4}, {4, 3}},
			want:       1,
		},

		{
			properties: [][]int{{2, 2}, {3, 3}},
			want:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
