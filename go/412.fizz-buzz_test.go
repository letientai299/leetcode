package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{
			1,
			[]string{
				"1",
			},
		},

		{
			5,
			[]string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
