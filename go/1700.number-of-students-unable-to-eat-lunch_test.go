package main

import "testing"

func Test_countStudents(t *testing.T) {
	tests := []struct {
		name       string
		students   []int
		sandwiches []int
		want       int
	}{
		{students: []int{1, 1, 0, 0}, sandwiches: []int{0, 1, 0, 1}, want: 0},
		{students: []int{1, 1, 1, 0, 0, 1}, sandwiches: []int{1, 0, 0, 0, 1, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStudents(tt.students, tt.sandwiches); got != tt.want {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
