package main

import "testing"

func Test_canReach(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		start int
		want  bool
	}{
		{
			arr:   []int{3, 0, 2, 1, 2},
			start: 2,
			want:  false,
		},

		{
			arr:   []int{4, 2, 3, 0, 3, 1, 2},
			start: 5,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.arr, tt.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
