package main

import (
	"reflect"
	"testing"
)

func Test_sequentialDigits(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want []int
	}{
		{
			low:  1000,
			high: 13000,
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},

		{
			low:  50,
			high: 788,
			want: []int{56, 67, 78, 89, 123, 234, 345, 456, 567, 678},
		},

		{
			low:  10,
			high: 788,
			want: []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678},
		},

		{
			low:  10,
			high: 1000,
			want: []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789},
		},

		{low: 100, high: 300, want: []int{123, 234}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequentialDigits(tt.low, tt.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sequentialDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
