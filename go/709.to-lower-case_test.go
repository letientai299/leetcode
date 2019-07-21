package main

import (
	"fmt"
	"testing"
)

func Test_toLowerCase(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{str: "hello", want: "hello"},
		{str: "HELLO", want: "hello"},
		{str: "HellO", want: "hello"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.str,
		)
		t.Run(testName, func(t *testing.T) {
			got := toLowerCase(tt.str)
			if got != tt.want {
				t.Errorf("toLowerCase(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
