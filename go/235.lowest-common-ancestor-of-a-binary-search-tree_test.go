package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v",
			tt.root, tt.p, tt.q,
		)
		t.Run(testName, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.root, tt.p, tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor(%v, %v, %v) = %v, want %v", tt.root, tt.p, tt.q, got, tt.want)
			}
		})
	}
}
