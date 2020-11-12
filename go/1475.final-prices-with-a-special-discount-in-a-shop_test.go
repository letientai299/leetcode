package main

import (
	"reflect"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	tests := []struct {
		prices []int
		want   []int
	}{
		{
			prices: []int{8, 4, 6, 2, 3},
			want:   []int{4, 2, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := finalPrices(tt.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
