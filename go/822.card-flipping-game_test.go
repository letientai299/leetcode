package main

import "testing"

func Test_flipgame(t *testing.T) {
	tests := []struct {
		name   string
		fronts []int
		backs  []int
		want   int
	}{
		{
			fronts: []int{1, 3, 2, 1, 1, 4, 4, 2, 4, 3},
			backs:  []int{1, 2, 4, 4, 3, 2, 3, 2, 3, 3},
			want:   4,
		},

		{
			fronts: []int{1, 2},
			backs:  []int{2, 1},
			want:   1,
		},

		{
			fronts: []int{1, 1},
			backs:  []int{2, 2},
			want:   1,
		},

		{
			fronts: []int{1, 2, 4, 4, 7},
			backs:  []int{1, 3, 4, 1, 3},
			want:   2,
		},

		{
			fronts: []int{1, 1},
			backs:  []int{1, 2},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipgame(tt.fronts, tt.backs); got != tt.want {
				t.Errorf("flipgame() = %v, want %v", got, tt.want)
			}
		})
	}
}
