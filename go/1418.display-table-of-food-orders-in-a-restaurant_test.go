package main

import (
	"reflect"
	"testing"
)

func Test_displayTable(t *testing.T) {
	tests := []struct {
		name   string
		orders [][]string
		want   [][]string
	}{
		{
			orders: [][]string{
				{"James", "12", "Fried Chicken"},
				{"Ratesh", "12", "Fried Chicken"},
				{"Amadeus", "12", "Fried Chicken"},
				{"Adam", "1", "Canadian Waffles"},
				{"Brianna", "1", "Canadian Waffles"},
			},
			want: [][]string{
				{"Table", "Canadian Waffles", "Fried Chicken"},
				{"1", "2", "0"},
				{"12", "0", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayTable(tt.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
