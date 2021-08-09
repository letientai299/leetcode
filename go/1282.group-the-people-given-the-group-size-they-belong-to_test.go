package main

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	tests := []struct {
		name       string
		groupSizes []int
		want       [][]int
	}{
		{
			groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
			want: [][]int{
				{0, 1, 2},
				{3, 4, 6},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
