package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	tests := []struct {
		area int
		want []int
	}{
		{4, []int{2, 2}},
		{1, []int{1, 1}},
		{2, []int{2, 1}},
		{3, []int{3, 1}},
		{5, []int{5, 1}},
		{6, []int{3, 2}},
		{7, []int{7, 1}},
		{8, []int{4, 2}},
		{9, []int{3, 3}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.area,
		)
		t.Run(testName, func(t *testing.T) {
			if got := constructRectangle(tt.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle(%v) = %v, want %v", tt.area, got, tt.want)
			}
		})
	}
}
