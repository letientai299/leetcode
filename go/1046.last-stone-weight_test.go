package main

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
