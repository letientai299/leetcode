package main

import "testing"

func Test_minFlips1318(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 5, b: 2, c: 8}, want: 4},
		{args: args{a: 4, b: 2, c: 7}, want: 1},
		{args: args{a: 1, b: 2, c: 3}, want: 0},
		{args: args{a: 2, b: 6, c: 5}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips1318(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("minFlips1318() = %v, want %v", got, tt.want)
			}
		})
	}
}
