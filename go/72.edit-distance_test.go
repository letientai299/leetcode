package main

import "testing"

func Test_minDistance72(t *testing.T) {
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
			if got := minDistance72(tt.a, tt.b); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
