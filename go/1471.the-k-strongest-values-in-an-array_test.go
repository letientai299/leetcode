package main

import (
	"reflect"
	"testing"
)

func Test_getStrongest(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{arr: []int{-7, 22, 17, 3}, k: 2, want: []int{22, 17}},
		{arr: []int{1, 1, 3, 5, 5}, k: 2, want: []int{5, 5}},
		{arr: []int{6, 7, 11, 7, 6, 8}, k: 5, want: []int{11, 8, 6, 6, 7}},
		{arr: []int{6, -3, 7, 2, 11}, k: 3, want: []int{-3, 11, 2}},
		{arr: []int{1, 2, 3, 4, 5}, k: 2, want: []int{5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStrongest(tt.arr, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStrongest() = %v, want %v", got, tt.want)
			}
		})
	}
}
