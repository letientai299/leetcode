package main

import (
	"reflect"
	"testing"
)

func Test_memLeak(t *testing.T) {
	tests := []struct {
		name    string
		memory1 int
		memory2 int
		want    []int
	}{
		{memory1: 2, memory2: 2, want: []int{3, 1, 0}},
		{memory1: 8, memory2: 11, want: []int{6, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memLeak(tt.memory1, tt.memory2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memLeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
