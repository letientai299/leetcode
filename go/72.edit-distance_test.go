package main

import "testing"

func Test_minDistance(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{
			a:    "sea",
			b:    "ate",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.a+":"+tt.b, func(t *testing.T) {
			if got := minDistance(tt.a, tt.b); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
