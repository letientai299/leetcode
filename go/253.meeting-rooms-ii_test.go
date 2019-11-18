package main

import (
	"fmt"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{{7, 10}, {2, 3}},
			want:      1,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.intervals,
		)
		t.Run(testName, func(t *testing.T) {
			got := minMeetingRooms(tt.intervals)
			if got != tt.want {
				t.Errorf("minMeetingRooms(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
