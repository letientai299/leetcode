package main

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	tests := []struct {
		deck []int
		want bool
	}{
		{
			deck: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: true,
		},
		{
			deck: []int{1, 1, 1, 2, 2, 2, 3, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
