package main

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 3, 2, 3, 1},
			want: []int{1, 3, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := frequencySort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
