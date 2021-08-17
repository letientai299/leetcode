package main

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	tests := []struct {
		name    string
		points  [][]int
		queries [][]int
		want    []int
	}{
		{
			points: [][]int{
				{99, 113}, {150, 165}, {23, 65}, {175, 154}, {84, 83}, {24, 59}, {124, 29}, {19, 97}, {117, 182}, {105, 191}, {83, 117}, {114, 35}, {0, 111}, {22, 53},
			},
			queries: [][]int{
				{105, 191, 155}, {114, 35, 94}, {84, 83, 68}, {175, 154, 28}, {99, 113, 80}, {175, 154, 177}, {175, 154, 181}, {114, 35, 134}, {22, 53, 105}, {124, 29, 164}, {6, 99, 39}, {84, 83, 35},
			},
			want: []int{
				11, 7, 8, 2, 7, 11, 13, 10, 10, 14, 3, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.points, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
