package main

import (
	"fmt"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{
			1, "A",
		},
		{
			2, "B",
		},
		{
			701, "ZY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := convertToTitle(tt.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}

	for i := 1; i <= 100; i++ {
		fmt.Printf("%8d | %8s\n", i, convertToTitle(i))
	}
}
