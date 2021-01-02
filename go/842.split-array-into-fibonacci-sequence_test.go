package main

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	tests := []struct {
		num  string
		want []int
	}{
		{num: "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511", want: nil},
		{num: "17522", want: []int{17, 5, 22}},
		{num: "123456579", want: []int{123, 456, 579}},
		{num: "11235813", want: []int{1, 1, 2, 3, 5, 8, 13}},
		{num: "011235813", want: []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{num: "112358130", want: nil},
		{num: "", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
