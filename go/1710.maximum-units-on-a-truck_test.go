package main

import "testing"

func Test_maximumUnits(t *testing.T) {
	tests := []struct {
		name      string
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{
			boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			truckSize: 10,
			want:      91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.boxTypes, tt.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
