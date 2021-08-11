package main

import "testing"

func Test_minSetSize(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			arr:  []int{11, 1, 1, 0},
			want: 1,
		},

		{
			arr:  []int{1},
			want: 1,
		},

		{
			arr:  []int{0, 1},
			want: 1,
		},

		{
			arr:  []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSetSize(tt.arr); got != tt.want {
				t.Errorf("minSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
