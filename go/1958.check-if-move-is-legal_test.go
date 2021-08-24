package main

import "testing"

func Test_checkMove(t *testing.T) {
	type args struct {
		board [][]byte
		rMove int
		cMove int
		color byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: [][]byte{
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', 'B', '.', '.', 'W', '.', '.', '.'},
					{'.', '.', 'W', '.', '.', '.', '.', '.'},
					{'.', '.', '.', 'W', 'B', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', 'B', 'W', '.', '.'},
					{'.', '.', '.', '.', '.', '.', 'W', '.'},
					{'.', '.', '.', '.', '.', '.', '.', 'B'},
				},
				rMove: 4,
				cMove: 4,
				color: 'W',
			},
			want: false,
		},
		{
			args: args{
				board: [][]byte{
					{'.', '.', 'W', '.', 'B', 'W', 'W', 'B'},
					{'B', 'W', '.', 'W', '.', 'W', 'B', 'B'},
					{'.', 'W', 'B', 'W', 'W', '.', 'W', 'W'},
					{'W', 'W', '.', 'W', '.', '.', 'B', 'B'},
					{'B', 'W', 'B', 'B', 'W', 'W', 'B', '.'},
					{'W', '.', 'W', '.', '.', 'B', 'W', 'W'},
					{'B', '.', 'B', 'B', '.', '.', 'B', 'B'},
					{'.', 'W', '.', 'W', '.', 'W', '.', 'W'},
				},
				rMove: 5,
				cMove: 4,
				color: 'W',
			},
			want: true,
		},

		{
			args: args{
				board: [][]byte{
					{'.', '.', '.', 'B', '.', '.', '.', '.'},
					{'.', '.', '.', 'W', '.', '.', '.', '.'},
					{'.', '.', '.', 'W', '.', '.', '.', '.'},
					{'.', '.', '.', 'W', '.', '.', '.', '.'},
					{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
					{'.', '.', '.', 'B', '.', '.', '.', '.'},
					{'.', '.', '.', 'B', '.', '.', '.', '.'},
					{'.', '.', '.', 'W', '.', '.', '.', '.'},
				},
				rMove: 4,
				cMove: 3,
				color: 'B',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMove(tt.args.board, tt.args.rMove, tt.args.cMove, tt.args.color); got != tt.want {
				t.Errorf("checkMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
