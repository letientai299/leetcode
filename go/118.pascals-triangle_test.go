package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{
			0,
			nil,
		},
		{
			2,
			[][]int{
				{1},
				{1, 1},
			},
		},
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			1,
			[][]int{
				{1},
			},
		},
		{
			3,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			6,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
