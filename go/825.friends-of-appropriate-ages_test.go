package main

import (
	"testing"
)

func Test_numFriendRequests(t *testing.T) {
	tests := []struct {
		name string
		ages []int
		want int
	}{
		{ages: []int{118, 14, 7, 63, 103}, want: 2},
		{ages: []int{73, 106, 39, 6, 26, 15, 30, 100, 71, 35, 46, 112, 6, 60, 110}, want: 29},
		{ages: []int{20, 30, 100, 100, 100, 110, 120, 120, 120, 120}, want: 37},
		{ages: []int{20, 30, 100, 110, 120}, want: 3},
		{ages: []int{16, 17, 18}, want: 2},
		{ages: []int{16, 16}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFriendRequests(tt.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
