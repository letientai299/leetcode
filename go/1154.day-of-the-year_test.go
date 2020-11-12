package main

import "testing"

func Test_dayOfYear(t *testing.T) {
	tests := []struct {
		date string
		want int
	}{
		{
			date: "2019-01-09",
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := dayOfYear(tt.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
