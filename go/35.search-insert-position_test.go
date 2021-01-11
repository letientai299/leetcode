package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]int{1, 3, 5, 6}, 5},
			want: 2,
		},

		{
			args: args{[]int{1, 3, 5, 6}, 2},
			want: 1},

		{
			args: args{[]int{1, 3, 5, 6}, 7},
			want: 4,
		},
		{
			args: args{[]int{1, 3, 5, 6}, 0},
			want: 0,
		},
		{
			args: args{[]int{1, 3, 5, 6}, 7},
			want: 4,
		},
		{
			args: args{[]int{}, 7},
			want: 0,
		},
		{
			args: args{[]int{7}, 7},
			want: 0,
		},
		{
			args: args{[]int{-1, 3, 5, 6}, -7},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
