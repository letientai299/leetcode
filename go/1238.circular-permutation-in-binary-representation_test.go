package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_circularPermutation(t *testing.T) {
	check := func(t *testing.T, arr []int) bool {
		for i := 1; i < len(arr); i++ {
			x := arr[i] & arr[i-1]
			if x&x-1 != 0 {
				t.Errorf("not diff only 1 bit, %d (%b), %d (%b)", arr[i-1], arr[i-1], arr[i], arr[i])
				return false
			}
		}

		n := len(arr) - 1
		x := arr[n] & arr[0]
		if x&x-1 != 0 {
			t.Errorf("not diff only 1 bit, %d (%b), %d (%b)", arr[0], arr[0], arr[n], arr[n])
			return false
		}
		return true
	}

	tests := []struct {
		n     int
		start int
	}{
		{n: 4, start: 0},
		{n: 4, start: 1},
		{n: 4, start: 3},
		{n: 2, start: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := circularPermutation(tt.n, tt.start)
			want := 1 << tt.n
			if len(got) != want {
				t.Errorf("invalid len, want=%v, got=%v", want, len(got))
				return
			}

			if assert.Equal(t, got[0], tt.start, "must start with %d", tt.start) {
				return
			}

			check(t, got)
		})
	}
}
