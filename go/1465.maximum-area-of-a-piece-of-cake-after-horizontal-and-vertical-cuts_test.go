package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				h: 5, w: 4,
				horizontalCuts: []int{3, 1},
				verticalCuts:   []int{1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.h, tt.args.w, tt.args.horizontalCuts, tt.args.verticalCuts); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
