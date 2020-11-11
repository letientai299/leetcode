package main

import "testing"

func Test_isMajorityElement(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{
			nums: []int{0, 1, 2, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isMajorityElement(tt.nums, tt.target); got != tt.want {
				t.Errorf("isMajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
