package main

import (
	"fmt"
	"testing"
)

func Test_canAttendMeetings(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      bool
	}{
		{
			intervals: [][]int{{0, 1}},
			want:      true,
		},

		{
			intervals: [][]int{{0, 1}, {1, 2}},
			want:      true,
		},

		{
			intervals: [][]int{{0, 2}, {1, 2}},
			want:      false,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.intervals,
		)
		t.Run(testName, func(t *testing.T) {
			got := canAttendMeetings(tt.intervals)
			if got != tt.want {
				t.Errorf("canAttendMeetings(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
