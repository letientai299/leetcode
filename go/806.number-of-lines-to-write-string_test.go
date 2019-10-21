package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	tests := []struct {
		widths []int
		S      string
		want   []int
	}{
		{
			widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			S:      "bbbcccdddaaa",
			want:   []int{2, 4},
		},

		{
			widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			S:      "abcdefghijklmnopqrstuvwxyz",
			want:   []int{3, 60},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.widths, tt.S,
		)
		t.Run(testName, func(t *testing.T) {
			if got := numberOfLines(tt.widths, tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines(%v, %v) = %v, want %v", tt.widths, tt.S, got, tt.want)
			}
		})
	}
}
