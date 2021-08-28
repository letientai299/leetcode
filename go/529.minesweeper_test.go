package main

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		click []int
		want  [][]byte
	}{
		{
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{3, 0},
			want: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.board, tt.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
