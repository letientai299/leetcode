package main

import (
	"reflect"
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	tests := []struct {
		seq  string
		want []int
	}{
		{seq: "(()())", want: []int{0, 1, 1, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.seq, func(t *testing.T) {
			if got := maxDepthAfterSplit(tt.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthAfterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
