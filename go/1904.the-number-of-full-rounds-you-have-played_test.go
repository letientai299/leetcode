package main

import "testing"

func Test_numberOfRounds(t *testing.T) {
	tests := []struct {
		name       string
		startTime  string
		finishTime string
		want       int
	}{
		{
			startTime:  "20:00",
			finishTime: "06:00",
			want:       40,
		},

		{
			startTime:  "12:01",
			finishTime: "12:44",
			want:       1,
		},

		{
			startTime:  "12:01",
			finishTime: "12:10",
			want:       0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRounds(tt.startTime, tt.finishTime); got != tt.want {
				t.Errorf("numberOfRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
