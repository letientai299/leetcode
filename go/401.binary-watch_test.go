package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readBinaryWatch(t *testing.T) {
	tests := []struct {
		num  int
		want []string
	}{
		{
			num:  1,
			want: []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"},
		},

		{
			num: 2,
			want: []string{
				"0:03", "0:05", "0:06", "0:09", "0:10", "0:12", "0:17", "0:18", "0:20",
				"0:24", "0:33", "0:34", "0:36", "0:40", "0:48",

				"1:01", "1:02", "1:04", "1:08", "1:16", "1:32",
				"2:01", "2:02", "2:04", "2:08", "2:16", "2:32",
				"3:00",
				"4:01", "4:02", "4:04", "4:08", "4:16", "4:32",
				"5:00",
				"6:00",
				"8:01", "8:02", "8:04", "8:08", "8:16", "8:32",
				"9:00",
				"10:00",
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := readBinaryWatch(tt.num)
			sort.Strings(got)
			assert.ElementsMatchf(t, got, tt.want, "got=%v", got)
		})
	}
}

func Test_genNum(t *testing.T) {
	tests := []struct {
		n    int
		slot uint
		want []int
	}{
		{n: 1, slot: 4, want: []int{1, 2, 4, 8}},
		{n: 2, slot: 4, want: []int{3, 5, 9, 6, 10, 12}},
		{n: 3, slot: 4, want: []int{7, 11, 13, 14}},
		{n: 4, slot: 4, want: []int{15}},
	}
	for _, tc := range tests {
		tt := tc
		got := genNum(tt.n, tt.slot)
		assert.ElementsMatch(t, got, tt.want, got)
	}
}
