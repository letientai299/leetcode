package main

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	tests := []struct {
		name  string
		rains []int
		want  []int
	}{
		{rains: []int{1, 0, 2, 3, 0, 1, 2}, want: []int{-1, 1, -1, -1, 2, -1, -1}},
		{rains: []int{0, 1, 1}, want: nil},
		{rains: []int{1, 0, 2, 0, 2, 1}, want: []int{-1, 1, -1, 2, -1, -1}},
		{rains: []int{69, 0, 0, 69}, want: []int{-1, 69, 1, -1}},
		{rains: []int{1, 2, 0, 0, 2, 1}, want: []int{-1, -1, 2, 1, -1, -1}},
		{rains: []int{1, 2, 3, 4}, want: []int{-1, -1, -1, -1}},
		{rains: []int{1, 2, 0, 1, 2}, want: nil},
		{rains: []int{10, 20, 20}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avoidFlood(tt.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
