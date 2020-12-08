package main

import "testing"

func Test_largestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{nums: []int{1, 0, 0, 0}, want: "1000"},
		{nums: []int{0, 0, 0}, want: "0"},
		{nums: []int{432, 43243}, want: "43243432"},
		{nums: []int{9, 1, 90, 91}, want: "991901"},
		{nums: []int{3, 30, 34, 5, 9}, want: "9534330"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
