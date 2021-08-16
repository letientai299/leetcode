package main

import "testing"

func Test_maximumNumber(t *testing.T) {
	tests := []struct {
		num    string
		change []int
		want   string
	}{
		{
			num:    "334111",
			change: []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5},
			want:   "334999",
		},

		{
			num:    "214010",
			change: []int{6, 7, 9, 7, 4, 0, 3, 4, 4, 7},
			want:   "974676",
		},

		{
			num:    "3",
			change: []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8},
			want:   "3",
		},

		{
			num:    "1232",
			change: []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8},
			want:   "8532",
		},

		{
			num:    "132",
			change: []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8},
			want:   "832",
		},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := maximumNumber(tt.num, tt.change); got != tt.want {
				t.Errorf("maximumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
