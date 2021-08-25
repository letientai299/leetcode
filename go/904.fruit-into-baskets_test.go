package main

import "testing"

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			fruits: []int{0, 0, 1, 1},
			want:   4,
		},

		{
			fruits: []int{1},
			want:   1,
		},

		{
			fruits: []int{1, 1},
			want:   2,
		},

		{
			fruits: []int{1, 2, 1},
			want:   3,
		},

		{
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},

		{
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
