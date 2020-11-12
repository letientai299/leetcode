package main

import "testing"

func Test_average(t *testing.T) {
	tests := []struct {
		salary []int
		want   float64
	}{
		{
			salary: []int{4, 3, 2, 1},
			want:   2.5,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := average(tt.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
