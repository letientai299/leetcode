package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	tests := []struct {
		left  int
		right int
		want  []int
	}{
		{
			left:  1,
			right: 22,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},

		{
			left:  1,
			right: 10,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.left, tt.right,
		)
		t.Run(testName, func(t *testing.T) {
			if got := selfDividingNumbers(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfDividingNumbers(%v, %v) = %v, want %v", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
