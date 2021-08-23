package main

import (
	"reflect"
	"testing"
)

func Test_processQueries(t *testing.T) {
	tests := []struct {
		name    string
		queries []int
		m       int
		want    []int
	}{
		{queries: []int{3, 1, 2, 1}, m: 5, want: []int{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processQueries(tt.queries, tt.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
