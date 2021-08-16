package main

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  []int
	}{
		{
			graph: [][]int{
				{1, 2, 3, 4}, {1, 2, 3, 4}, {3, 4}, {4}, {},
			},
			want: []int{2, 3, 4},
		},

		{
			graph: [][]int{{1, 3, 4, 5, 7, 9}, {1, 3, 8, 9}, {3, 4, 5, 8}, {1, 8}, {5, 7, 8}, {8, 9}, {7, 8, 9}, {3}, {}, {}},
			want:  []int{5, 8, 9},
		},
		{
			graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}},
			want:  []int{4},
		},
		{
			graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}},
			want:  []int{2, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
