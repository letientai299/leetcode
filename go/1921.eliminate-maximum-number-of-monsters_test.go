package main

import "testing"

func Test_eliminateMaximum(t *testing.T) {
	tests := []struct {
		name  string
		dist  []int
		speed []int
		want  int
	}{
		{dist: []int{3, 2, 4}, speed: []int{5, 3, 2}, want: 1},
		{dist: []int{1, 1, 2, 3}, speed: []int{1, 1, 1, 1}, want: 1},
		{dist: []int{1, 3, 4}, speed: []int{1, 1, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.dist, tt.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
