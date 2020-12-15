package main

import (
	"reflect"
	"testing"
)

func Test_getSumAbsoluteDifferences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: []int{10, 7, 6, 7, 10}},
		{nums: []int{2, 3, 5}, want: []int{4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumAbsoluteDifferences(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
