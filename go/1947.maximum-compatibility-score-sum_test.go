package main

import "testing"

func Test_maxCompatibilitySum(t *testing.T) {
	tests := []struct {
		name     string
		students [][]int
		mentors  [][]int
		want     int
	}{
		{
			students: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}},
			mentors:  [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}},
			want:     8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCompatibilitySum(tt.students, tt.mentors); got != tt.want {
				t.Errorf("maxCompatibilitySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
