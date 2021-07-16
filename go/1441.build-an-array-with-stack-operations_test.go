package main

import (
	"reflect"
	"testing"
)

func Test_buildArray_1441(t *testing.T) {
	tests := []struct {
		target []int
		n      int
		want   []string
	}{
		{target: []int{2, 3, 4}, n: 3, want: []string{"Push", "Pop", "Push", "Push", "Push"}},
		{target: []int{2}, n: 3, want: []string{"Push", "Pop", "Push"}},
		{target: []int{1, 2}, n: 3, want: []string{"Push", "Push"}},
		{target: []int{1, 3}, n: 3, want: []string{"Push", "Push", "Pop", "Push"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := buildarray1441(tt.target, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
