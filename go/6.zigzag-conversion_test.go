package main

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		want    string
	}{
		{s: "abcde", numRows: 4, want: "abced"},
		{s: "paypalishiring", numRows: 4, want: "pinalsigyahrpi"},
		{s: "paypalishiring", numRows: 3, want: "pahnaplsiigyir"},
		{s: "paypalishiring", numRows: 40, want: "paypalishiring"},
		{s: "paypalishiring", numRows: 1, want: "paypalishiring"},
		{s: "paypalishiring", numRows: 2, want: "pyaihrnaplsiig"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.numRows,
		)
		t.Run(testName, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			if got != tt.want {
				t.Errorf("convert(%v, %v) = %v, want %v", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
