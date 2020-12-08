package main

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		t    int
		want bool
	}{
		{nums: []int{1, 2, 3, 1}, k: 3, t: 0, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
