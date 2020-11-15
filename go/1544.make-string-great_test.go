package main

import "testing"

func Test_makeGood(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "aA",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := makeGood(tt.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
