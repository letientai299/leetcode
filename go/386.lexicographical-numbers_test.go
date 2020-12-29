package main

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lexicalOrder(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 19},
		{n: 1229},
		{n: 14959},
		{n: 192},
		{n: 23},
		{n: 13},
		{n: 10},
		{n: 100},
		{n: 2000},
		{n: 325},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := lexicalOrder(tt.n)
			want := lexicalOrderBruteForce(tt.n)
			assert.Equal(t, want, got)
		})
	}
}

func lexicalOrderBruteForce(n int) []int {
	arr := make([]int, n)
	ss := make([]string, n)
	for i := range ss {
		ss[i] = strconv.Itoa(i + 1)
	}
	sort.Strings(ss)
	for i := range arr {
		arr[i], _ = strconv.Atoi(ss[i])
	}
	return arr
}
