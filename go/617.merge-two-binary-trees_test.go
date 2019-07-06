package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	tests := []struct {
		t1   *TreeNode
		t2   *TreeNode
		want *TreeNode
	}{

		{
			t1:   treeFromLevelOrder(),
			t2:   treeFromLevelOrder(1, 2, 3),
			want: treeFromLevelOrder(1, 2, 3),
		},

		{
			t1:   treeFromLevelOrder(1, 2, 3),
			t2:   treeFromLevelOrder(),
			want: treeFromLevelOrder(1, 2, 3),
		},

		{
			t1:   treeFromLevelOrder(1, 3, 2, 5),
			t2:   treeFromLevelOrder(2, 1, 3, NA, 4, NA, 7),
			want: treeFromLevelOrder(3, 4, 5, 5, 4, NA, 7),
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.t1, tt.t2,
		)
		t.Run(testName, func(t *testing.T) {
			if got := mergeTrees(tt.t1.Clone(), tt.t2.Clone()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees(%v, %v) = %v, want %v", tt.t1, tt.t2, got, tt.want)
			}
		})
	}
}
