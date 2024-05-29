package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCommon(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getCommon(tt.args.nums1, tt.args.nums2), "getCommon(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
