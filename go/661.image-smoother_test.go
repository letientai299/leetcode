package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	tests := []struct {
		img  [][]int
		want [][]int
	}{
		{
			img: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.img,
		)
		t.Run(testName, func(t *testing.T) {
			if got := imageSmoother(tt.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother(%v) = %v, want %v", tt.img, got, tt.want)
			}
		})
	}
}
