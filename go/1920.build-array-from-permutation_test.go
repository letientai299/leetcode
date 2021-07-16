package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			nums: []int{0, 2, 1, 5, 3, 4},
			want: []int{0, 1, 2, 4, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
