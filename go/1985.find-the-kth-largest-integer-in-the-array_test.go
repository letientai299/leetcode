package main

import "testing"

func Test_kthLargestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []string
		k    int
		want string
	}{
		{
			nums: []string{"3", "6", "7", "10"},
			k:    4,
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestNumber(tt.nums, tt.k); got != tt.want {
				t.Errorf("kthLargestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
