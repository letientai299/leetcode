package main

import (
	"reflect"
	"testing"
)

func Test_mostVisited(t *testing.T) {
	tests := []struct {
		n      int
		rounds []int
		want   []int
	}{
		{
			n:      2,
			rounds: []int{2, 1, 2, 1, 2, 1, 2, 1, 2},
			want:   []int{2},
		},
		{
			n:      4,
			rounds: []int{1, 3, 1, 2},
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := mostVisited(tt.n, tt.rounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}
