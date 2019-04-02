package main

import (
	"fmt"
	"testing"
)

func Test_toHex(t *testing.T) {
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
		{7, "7"},
		{8, "8"},
		{9, "9"},
		{10, "a"},
		{11, "b"},
		{-1, "ffffffff"},
		{-2, "fffffffe"},
		{-3, "fffffffd"},
		{15, "f"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := toHex(tt.num)
			if got != tt.want {
				t.Errorf("toHex(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
