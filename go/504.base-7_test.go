package main

import (
	"fmt"
	"testing"
)

func Test_convertToBase7(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{3, "3"},
		{4, "4"},
		{5, "5"},
		{6, "6"},
		{-1, "-1"},
		{-2, "-2"},
		{-3, "-3"},
		{-4, "-4"},
		{-5, "-5"},
		{-6, "-6"},
		{7, "10"},
		{8, "11"},
		{9, "12"},
		{10, "13"},
		{-7, "-10"},
		{100, "202"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := convertToBase7(tt.num)
			if got != tt.want {
				t.Errorf("convertToBase7(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
