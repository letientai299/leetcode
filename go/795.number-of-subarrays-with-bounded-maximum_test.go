package main

import "testing"

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		a []int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{a: []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, l: 32, r: 69},
			want: 22,
		},

		{
			args: args{a: []int{2, 1, 4, 3}, l: 2, r: 3},
			want: 3,
		},

		{
			args: args{a: []int{2, 9, 2, 5, 6}, l: 2, r: 8},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayBoundedMax(tt.args.a, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
