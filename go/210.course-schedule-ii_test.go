package main

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			numCourses:    3,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 2, 1},
		},


		{
			numCourses:    2,
			prerequisites: [][]int{},
			want:          []int{0, 1},
		},


		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 1, 2, 3},
		},

		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          []int{},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}},
			want:          []int{0, 1, 2, 3},
		},

		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.numCourses, tt.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
