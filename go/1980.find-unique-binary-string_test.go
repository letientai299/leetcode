package main

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	tests := []struct {
		name string
		nums []string
		want string
	}{
		{
			nums: []string{"01", "00"},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.nums); got != tt.want {
				t.Errorf("findDifferentBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
