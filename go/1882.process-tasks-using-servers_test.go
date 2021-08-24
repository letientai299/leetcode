package main

import (
	"reflect"
	"testing"
)

func Test_assignTasks(t *testing.T) {
	tests := []struct {
		name    string
		servers []int
		tasks   []int
		want    []int
	}{
		{
			servers: []int{91, 74, 89, 69, 11, 56, 33, 94, 98, 51, 95, 51, 25, 37},
			tasks:   []int{83, 53, 10, 89, 64, 54, 44, 72, 38, 78,  4, 83, 21, 92, 31},
			want:    []int{4, 12, 6, 13, 9, 11, 5, 3, 1, 2, 0, 7, 6, 10, 0},
		},

		{
			servers: []int{5, 1, 4, 3, 2},
			tasks:   []int{2, 1, 2, 4, 5, 2, 1},
			want:    []int{1, 4, 1, 4, 1, 3, 2},
		},

		{
			servers: []int{3, 3, 2},
			tasks:   []int{1, 2, 3, 2, 1, 2},
			want:    []int{2, 2, 0, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignTasks(tt.servers, tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
