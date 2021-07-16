package main

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 1}, want: []int{1, 2, 1, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenation(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
