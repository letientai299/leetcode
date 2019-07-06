package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRestaurant(t *testing.T) {
	tests := []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{
			list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			want:  []string{"Shogun"},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.list1, tt.list2,
		)
		t.Run(testName, func(t *testing.T) {
			got := findRestaurant(tt.list1, tt.list2)
			assert.ElementsMatchf(t, got, tt.want, "expect %v, got %v", tt.want, got)
		})
	}
}
