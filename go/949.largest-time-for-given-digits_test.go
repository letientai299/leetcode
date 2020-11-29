package main

import "testing"

func Test_largestTimeFromDigits(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want string
	}{
		{arr: []int{3, 9, 2, 7}, want: ""},
		{arr: []int{0, 0, 0, 0}, want: "00:00"},
		{arr: []int{1, 0, 5, 7}, want: "17:50"},
		{arr: []int{1, 9, 5, 7}, want: "19:57"},
		{arr: []int{5, 5, 5, 5}, want: ""},
		{arr: []int{1, 2, 3, 4}, want: "23:41"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTimeFromDigits(tt.arr); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
