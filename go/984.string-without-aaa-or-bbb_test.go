package main

import "testing"

func Test_strWithout3a3b(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want string
	}{
		{a: 1, b: 2, want: "bab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strWithout3a3b(tt.a, tt.b); got != tt.want {
				t.Errorf("strWithout3a3b() = %v, want %v", got, tt.want)
			}
		})
	}
}
