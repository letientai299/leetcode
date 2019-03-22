package main

import (
	"reflect"
	"testing"
)

func Test_twoSum_1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sample on leetcode",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "simple case",
			args: args{
				nums:   []int{0, 0, 1},
				target: 1,
			},
			want: []int{0, 2},
		},
		{
			name: "with negative",
			args: args{
				nums:   []int{2, 3, 4, 5, -1, 1},
				target: 0,
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum_1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
