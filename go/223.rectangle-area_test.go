package main

import "testing"

func Test_computeArea(t *testing.T) {
	tests := []struct {
		args [8]int
		want int
	}{
		{args: [8]int{-5, 0, 5, 5, -3, -3, 3, 3}, want: 68},

		{
			args: [8]int{-2, -2, 2, 2, 1, -3, 3, 3},
			want: 24,
		},

		{
			args: [8]int{-5, 0, 5, 3, -3, -3, 3, 3},
			want: 48,
		},

		{
			args: [8]int{-5, -2, 5, 1, -3, -3, 3, 3},
			want: 48,
		},

		{
			args: [8]int{-2, -2, 2, 2, -3, -3, 3, -1},
			want: 24,
		},

		{
			args: [8]int{-3, 0, 3, 4, 0, -1, 9, 2},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := computeArea(
				tt.args[0],
				tt.args[1],
				tt.args[2],
				tt.args[3],
				tt.args[4],
				tt.args[5],
				tt.args[6],
				tt.args[7],
			); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
