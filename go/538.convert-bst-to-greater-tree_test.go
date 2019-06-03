package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			if got := convertBST(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
