package main

import "testing"

func Test_addRungs(t *testing.T) {
	tests := []struct {
		name  string
		rungs []int
		dist  int
		want  int
	}{
		{
			rungs: []int{3, 6, 8, 10},
			dist:  3,
			want:  0,
		},

		{
			rungs: []int{1, 3, 5, 10},
			dist:  2,
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addRungs(tt.rungs, tt.dist); got != tt.want {
				t.Errorf("addRungs() = %v, want %v", got, tt.want)
			}
		})
	}
}
