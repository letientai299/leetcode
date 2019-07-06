package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []float64
	}{
		{
			root: treeFromLevelOrder(3, 9, 20, NA, 15, 7),
			want: []float64{3, 14.5, 11},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			if got := averageOfLevels(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
