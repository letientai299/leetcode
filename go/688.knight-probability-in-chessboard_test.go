package main

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{n: 8, k: 30, row: 6, column: 4},
			want: 0.0001905256629833365,
		},

		{
			args: args{n: 3, k: 2, row: 0, column: 0},
			want: 0.06250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
