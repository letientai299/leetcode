package main

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		S      string
		shifts []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{shifts: []int{10, 20, 30}, S: "bad"},
			want: "jyh",
		},

		{
			args: args{shifts: []int{3, 5, 9}, S: "abc"},
			want: "rpl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.S, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
