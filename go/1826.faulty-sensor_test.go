package main

import "testing"

func Test_badSensor(t *testing.T) {
	tests := []struct {
		name    string
		sensor1 []int
		sensor2 []int
		want    int
	}{
		{
			sensor1: []int{1, 2, 3, 2, 3, 2},
			sensor2: []int{1, 2, 3, 3, 2, 3},
			want:    -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := badSensor(tt.sensor1, tt.sensor2); got != tt.want {
				t.Errorf("badSensor() = %v, want %v", got, tt.want)
			}
		})
	}
}
