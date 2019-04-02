package main

import (
	"fmt"
	"testing"
)

func Test_addStrings(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "1",
			num2: "2",
			want: "3",
		},

		{
			num1: "0",
			num2: "2",
			want: "2",
		},

		{
			num1: "0",
			num2: "0",
			want: "0",
		},

		{
			num1: "123456",
			num2: "1234567",
			want: "1358023",
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.num1, tt.num2,
		)
		t.Run(testName, func(t *testing.T) {
			got := addStrings(tt.num1, tt.num2)
			if got != tt.want {
				t.Errorf("addStrings(%v, %v) = %v, want %v", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}
